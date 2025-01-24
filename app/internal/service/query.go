package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage"
	"io/ioutil"
	"log/slog"
	"net/http"
	"time"
)

type QueryService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *slog.Logger
}

func NewQueryService(repository *storage.Repository, config *config.Config, logger *slog.Logger) *QueryService {
	return &QueryService{
		repository: repository,
		config:     config,
		logger:     logger,
	}
}

func (s *QueryService) ProcessQuery(ctx context.Context, query string) (string, error) {
	s.logger.Info("Processing query", slog.String("query", query))

	cachedQuery, err := s.repository.Redis.Query.GetQueryByRequest(ctx, query)
	if err != nil {
		s.logger.Error("Failed to get query from Redis", slog.String("query", query), slog.Any("error", err))
		return "", err
	}

	if cachedQuery != nil {
		s.logger.Info("Query found in Redis cache", slog.String("query", query))
		return cachedQuery.Response, nil
	}

	response, err := s.processWithTinyLlama(ctx, query)
	if err != nil {
		s.logger.Error("Failed to process query with TinyLlama", slog.String("query", query), slog.Any("error", err))
		return "", err
	}
	err = s.repository.Redis.Query.CreateQuery(ctx, &model.Query{
		Text:      query,
		Response:  response.Response["response"].(string),
		CreatedAt: time.Now(),
	})
	if err != nil {
		s.logger.Error("Failed to save query to Redis", slog.String("query", query), slog.Any("error", err))
		return "", err
	}

	id, err := s.repository.Postgres.Query.CreateQuery(ctx, &model.Query{
		Text:      query,
		Response:  response.Response["response"].(string),
		CreatedAt: time.Now(),
	})
	if err != nil {
		s.logger.Error("Failed to save query to Postgres", slog.String("query", query), slog.Any("error", err))
		return "", err
	}
	response.ID = id
	response.CreatedAt = time.Now()
	err = s.repository.Mongo.Query.CreateQuery(ctx, response)
	if err != nil {
		s.logger.Error("Failed to save query to Mongo", slog.String("query", query), slog.Any("error", err))
		return "", err
	}

	s.logger.Info("Query processed successfully", slog.String("query", query))
	return response.Response["response"].(string), nil
}

func (s *QueryService) processWithTinyLlama(ctx context.Context, query string) (*model.QueryWithRawResponse, error) {
	url := s.config.Ollama.URL

	requestBody, err := json.Marshal(map[string]interface{}{
		"model":  "tinyllama",
		"prompt": query,
		"stream": false,
	})
	if err != nil {
		s.logger.Error("Failed to marshal request body", slog.Any("error", err))
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		s.logger.Error("Failed to create HTTP request", slog.Any("error", err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request to TinyLlama", slog.Any("error", err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		errMsg := fmt.Sprintf("TinyLlama API returned status: %d, response: %s", resp.StatusCode, string(body))
		s.logger.Error(errMsg)
		return nil, fmt.Errorf(errMsg)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body", slog.Any("error", err))
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		s.logger.Error("Failed to unmarshal response body into map", slog.Any("error", err))
		return nil, err
	}

	s.logger.Info("Successfully processed query with TinyLlama", slog.String("query", query), slog.String("response", result["response"].(string)))
	return &model.QueryWithRawResponse{
		Text:     query,
		Response: result,
	}, nil
}

func (s *QueryService) GetHistory(ctx context.Context) ([]model.Query, error) {
	s.logger.Info("Fetching query history")

	history, err := s.repository.Postgres.Query.GetAllQueries(ctx)
	if err != nil {
		s.logger.Error("Failed to fetch query history from Postgres", slog.Any("error", err))
		return nil, err
	}

	if len(history) == 0 {
		s.logger.Warn("Query history is empty")
		return nil, errors.New("history is empty")
	}

	s.logger.Info("Query history fetched successfully", slog.Int("history_count", len(history)))
	return history, nil
}
