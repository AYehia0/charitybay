package app

import (
	"charitybay/config"
	v1 "charitybay/internal/controller/http/v1"
	"charitybay/internal/usecase"
	"charitybay/internal/usecase/repo"
	"charitybay/pkg/db"
	httpserver "charitybay/pkg/http"
	"charitybay/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	logger := logger.New(cfg.Log.Level)

	// database
	pg, err := db.New(cfg.Postgres.Url, db.MaxPoolSize(cfg.Postgres.MaxPoolSize))

	if err != nil {
		logger.Fatalf("DB error: %s", err)
	}

	// usecases
	greetUsecase := usecase.New(repo.New(pg))

	// http server
	handler := gin.New()
	v1.NewRouter(handler, logger, greetUsecase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Http.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Infof("app - Run - signal: %s", s)
	case err = <-httpServer.Notify():
		logger.Errorf("app - Run - httpServer.Notify: %s", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Errorf("app - Run - httpServer.Shutdown: %s", err)
	}
}
