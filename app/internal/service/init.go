package service

import (
	"context"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage"
	"log/slog"
)

type Service struct {
	Query IQueryService
	Auth  IAuthService
}

func New(repository *storage.Repository, config *config.Config, logger *slog.Logger) *Service {
	return &Service{
		Query: NewQueryService(repository, config, logger),
		Auth:  NewAuthService(repository, config, logger),
	}
}

type IQueryService interface {
	ProcessQuery(ctx context.Context, query string) (string, error)
	GetHistory(ctx context.Context) ([]model.Query, error)
}

type IAuthService interface {
	VerifyPassword(ctx context.Context, username, password string) (bool, error)
	Register(ctx context.Context, username, password string) (int, error)
}
