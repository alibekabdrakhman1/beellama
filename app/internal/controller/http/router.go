package http

import (
	_ "github.com/alibekabdrakhman1/beellama/app/internal/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func (s *Server) SetupRoutes() {
	s.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := s.App.Group("/api")
	api.POST("/register", s.handler.User.Register)

	wm := api.Group("")
	wm.Use(s.AuthMiddleware.ValidateAuth())
	wm.POST("/process", s.handler.Query.Process)
	wm.GET("/history", s.handler.Query.History)
}
