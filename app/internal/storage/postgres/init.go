package postgres

import (
	"context"
	"database/sql"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
)

type IQueryRepository interface {
	CreateQuery(ctx context.Context, query *model.Query) (int, error)
	GetAllQueries(ctx context.Context) ([]model.Query, error)
}

type IUserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type RepositoryPostgres struct {
	Query IQueryRepository
	User  IUserRepository
}

func New(db *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{
		Query: NewQueryRepository(db),
		User:  NewUserRepository(db),
	}
}
