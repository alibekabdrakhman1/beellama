package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/alibekabdrakhman1/beellama/app/internal/controller/http/middleware"
	"log"
	"net/http"
	"time"

	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/controller/http/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg            *config.Config
	handler        *handler.Manager
	App            *gin.Engine
	AuthMiddleware *middleware.BasicAuth
}

func NewServer(cfg *config.Config, handler *handler.Manager, auth *middleware.BasicAuth) *Server {
	return &Server{
		cfg:            cfg,
		handler:        handler,
		AuthMiddleware: auth,
	}
}

func (s *Server) StartHTTPServer(ctx context.Context) error {
	s.App = s.BuildEngine()
	s.SetupRoutes()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", s.cfg.Server.Port),
		Handler: s.App,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen:%v\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%v", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	return r
}
