package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"time"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

type UserRepository struct {
	DB *sql.DB
}

// Создает User-a и возвращает айдишку
func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var id int
	err := r.DB.QueryRowContext(ctx, `
		INSERT INTO users (username, password, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Username, user.Password, user.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Для сверки пароля
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var user model.User
	err := r.DB.QueryRowContext(ctx, `
		SELECT username, password
		FROM users
		WHERE username = $1
	`, username).Scan(&user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return user, nil
}
