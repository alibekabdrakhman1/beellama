package applicator

import (
	"context"
	"github.com/alibekabdrakhman1/beellama/app/internal/config"
	"github.com/alibekabdrakhman1/beellama/app/internal/controller/http"
	"github.com/alibekabdrakhman1/beellama/app/internal/controller/http/handler"
	"github.com/alibekabdrakhman1/beellama/app/internal/controller/http/middleware"
	"github.com/alibekabdrakhman1/beellama/app/internal/service"
	"github.com/alibekabdrakhman1/beellama/app/internal/storage"
	"github.com/alibekabdrakhman1/beellama/app/pkg/database"
	"log"
	"log/slog"
	"os"
	"os/signal"
)

type App struct {
	logger *slog.Logger
	config *config.Config
}

func New(logger *slog.Logger, cfg *config.Config) *App {
	return &App{
		config: cfg,
		logger: logger,
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	gracefullyShutdown(cancel)

	postgres, err := database.DialPostgres(a.config)
	if err != nil {
		return err
	}
	mongoClient, err := database.DialMongo(a.config)
	if err != nil {
		return err
	}
	redisClient, err := database.DialRedis(a.config)
	if err != nil {
		return err
	}

	repository := storage.New(postgres, mongoClient, redisClient)

	srv := service.New(repository, a.config, a.logger)

	endPointHandler := handler.New(srv, a.logger)

	basicAuthMiddleware := middleware.NewBasicAuth(srv.Auth, a.logger)
	HTTPServer := http.NewServer(a.config, endPointHandler, basicAuthMiddleware)
	return HTTPServer.StartHTTPServer(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
