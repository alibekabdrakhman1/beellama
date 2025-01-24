package service

import (
	"context"
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage"
	"log/slog"
)

type AuthService struct {
	repository *storage.Repository
	config     *config.Config
	logger     *slog.Logger
}

func NewAuthService(repository *storage.Repository, config *config.Config, logger *slog.Logger) *AuthService {
	return &AuthService{
		repository: repository,
		config:     config,
		logger:     logger,
	}
}

func (s *AuthService) VerifyPassword(ctx context.Context, username, password string) (bool, error) {
	user, err := s.repository.Postgres.User.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		return false, err
	}

	if match {
		return true, nil
	}
	return false, nil
}

func (s *AuthService) Register(ctx context.Context, username, password string) (int, error) {
	existingUser, _ := s.repository.Postgres.User.GetUserByUsername(ctx, username)
	if existingUser.Username != "" {
		return 0, errors.New("user already exists")
	}

	hashedPassword, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return 0, err
	}

	userID, err := s.repository.Postgres.User.CreateUser(ctx, &model.User{
		Username: username,
		Password: hashedPassword,
	})
	if err != nil {
		return 0, err
	}

	return userID, nil
}
