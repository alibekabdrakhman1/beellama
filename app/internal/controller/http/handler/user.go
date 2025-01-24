package handler

import (
	"github.com/alibekabdrakhman1/beellama/app/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUserHandler(service *service.Service, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

type UserHandler struct {
	service *service.Service
	logger  *slog.Logger
}

// Register @Summary User registration
// @Description Registers a new user.
// @ID user-register
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.User true "Registration request payload"
// @Success 200 {object} ResponseProcessSuccess "Successful registration response"
// @Failure 400 {object} ResponseProcessFailure "Bad Request"
// @Router /api/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var request struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.Auth.Register(c.Request.Context(), request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
