package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewQueryRepository(client *redis.Client) *QueryRepository {
	return &QueryRepository{
		Client: client,
	}
}

type QueryRepository struct {
	Client *redis.Client
}

func (r *QueryRepository) GetQueryByRequest(ctx context.Context, request string) (*model.Query, error) {
	data, err := r.Client.Get(ctx, request).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var query model.Query
	if err := json.Unmarshal([]byte(data), &query); err != nil {
		return nil, err
	}

	return &query, nil
}

func (r *QueryRepository) CreateQuery(ctx context.Context, query *model.Query) error {
	data, err := json.Marshal(query)
	if err != nil {
		return err
	}

	err = r.Client.Set(ctx, query.Text, data, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
