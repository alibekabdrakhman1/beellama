package mongo

import (
	"context"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type IQueryRepository interface {
	CreateQuery(ctx context.Context, query *model.QueryWithRawResponse) error
}

type RepositoryMongo struct {
	Query IQueryRepository
}

func New(client *mongo.Database) *RepositoryMongo {
	return &RepositoryMongo{
		Query: NewQueryRepository(client),
	}
}
