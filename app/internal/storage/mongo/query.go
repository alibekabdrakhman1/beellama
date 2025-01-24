package mongo

import (
	"context"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func NewQueryRepository(client *mongo.Database) *QueryRepository {
	return &QueryRepository{
		DB: client,
	}
}

type QueryRepository struct {
	DB *mongo.Database
}

func (r *QueryRepository) CreateQuery(ctx context.Context, query *model.QueryWithRawResponse) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	collection := r.DB.Collection("queries")

	_, err := collection.InsertOne(ctx, query)
	return err
}
