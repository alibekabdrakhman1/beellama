package postgres

import (
	"context"
	"database/sql"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"time"
)

func NewQueryRepository(db *sql.DB) *QueryRepository {
	return &QueryRepository{
		DB: db,
	}
}

type QueryRepository struct {
	DB *sql.DB
}

func (r *QueryRepository) GetAllQueries(ctx context.Context) ([]model.Query, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, text, response, created_at
		FROM queries
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var queries []model.Query
	for rows.Next() {
		var query model.Query
		if err := rows.Scan(&query.ID, &query.Text, &query.Response, &query.CreatedAt); err != nil {
			return nil, err
		}
		queries = append(queries, query)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return queries, nil
}

func (r *QueryRepository) CreateQuery(ctx context.Context, query *model.Query) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int
	err := r.DB.QueryRowContext(ctx, `
		INSERT INTO queries (text, response, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, query.Text, query.Response, query.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
