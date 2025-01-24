package database

import (
	"database/sql"
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func DialPostgres(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Postgres.Host, cfg.Database.Postgres.Port, cfg.Database.Postgres.User, cfg.Database.Postgres.Password, cfg.Database.Postgres.DBName, cfg.Database.Postgres.SSLMode)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при пинге postgres: %v", err)
	}

	log.Println("подключено к postgres")
	return db, nil
}
