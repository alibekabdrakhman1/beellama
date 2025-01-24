package middleware

import (
	"github.com/alibekabdrakhman1/beellama/app/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type BasicAuth struct {
	AuthService service.IAuthService
	logger      *slog.Logger
}

func NewBasicAuth(authService service.IAuthService, logger *slog.Logger) *BasicAuth {
	return &BasicAuth{
		AuthService: authService,
		logger:      logger,
	}
}

func (m *BasicAuth) ValidateAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Отсутствует заголовок авторизации",
			})
			c.Abort()
			return
		}

		ok, err := m.AuthService.VerifyPassword(c.Request.Context(), user, pass)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Ошибка при проверке данных авторизации",
			})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Неверное имя пользователя или парольd",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
