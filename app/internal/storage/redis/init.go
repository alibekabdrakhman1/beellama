package redis

import (
	"context"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/redis/go-redis/v9"
)

type IQueryRepository interface {
	CreateQuery(ctx context.Context, query *model.Query) error
	GetQueryByRequest(ctx context.Context, request string) (*model.Query, error)
}

type RepositoryRedis struct {
	Query IQueryRepository
}

func New(client *redis.Client) *RepositoryRedis {
	return &RepositoryRedis{
		Query: NewQueryRepository(client),
	}
}
