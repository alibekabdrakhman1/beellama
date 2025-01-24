package database

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func DialRedis(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ошибка при пинге redis: %v", err)
	}

	log.Println("подключено к редису")
	return client, nil
}
