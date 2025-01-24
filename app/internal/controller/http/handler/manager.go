package handler

import (
	"github.com/alibekabdrakhman1/beellama/app/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func New(service *service.Service, logger *slog.Logger) *Manager {
	return &Manager{
		User:  NewUserHandler(service, logger),
		Query: NewQueryHandler(service, logger),
	}
}

type Manager struct {
	User  IUserHandler
	Query IQueryHandler
}

type IUserHandler interface {
	Register(c *gin.Context)
}

type IQueryHandler interface {
	Process(c *gin.Context)
	History(c *gin.Context)
}
