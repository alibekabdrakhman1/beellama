package storage

import (
	"database/sql"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage/mongo"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage/postgres"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage/redis"
	redis2 "github.com/redis/go-redis/v9"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Mongo    *mongo.RepositoryMongo
	Postgres *postgres.RepositoryPostgres
	Redis    *redis.RepositoryRedis
}

func New(db *sql.DB, mongoClient *mongo2.Database, redisClient *redis2.Client) *Repository {
	return &Repository{
		Mongo:    mongo.New(mongoClient),
		Postgres: postgres.New(db),
		Redis:    redis.New(redisClient),
	}
}
