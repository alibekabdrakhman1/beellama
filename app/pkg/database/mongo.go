package database

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func DialMongo(cfg *config.Config) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Database.Mongo.URI))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании клиента mongodb: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("ошибка при подключении к mongodb: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ошибка при пинге mongodb: %v", err)
	}

	log.Println("подключено к mongodb!")
	return client.Database(cfg.Database.Mongo.Database), nil
}
