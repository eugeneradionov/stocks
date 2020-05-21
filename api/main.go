package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/handlers"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/services"
	"github.com/eugeneradionov/stocks/api/store/repo"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/eugeneradionov/stocks/api/transport/rabbitmq"
	"github.com/eugeneradionov/stocks/api/validator"
	"go.uber.org/zap"
)

const defaultConfigPath = "config.json"

// nolint
func main() {
	err := config.Load(defaultConfigPath)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	err = logger.Load(config.Get())
	if err != nil {
		log.Fatalf("Failed to laod logger: %s", err.Error())
	}

	err = validator.Load()
	if err != nil {
		logger.Get().Error("Failed to load validator", zap.Error(err))
	}

	err = postgres.Load(config.Get().Postgres, logger.Get().Sugar())
	if err != nil {
		logger.Get().Fatal("Failed to connect to postgres", zap.Error(err))
	}

	err = rabbitmq.Load(config.Get().RabbitMQ)
	if err != nil {
		logger.Get().Fatal("Failed to connect to rabbit", zap.Error(err))
	}

	err = repo.Load()
	if err != nil {
		logger.Get().Fatal("Failed to initialize postgres repo", zap.Error(err))
	}

	err = services.Load(config.Get(), rabbitmq.Get())
	if err != nil {
		logger.Get().Fatal("Failed to load services", zap.Error(err))
	}

	extErr := services.Get().Stocks().ConsumeAll()
	if extErr != nil {
		logger.Get().Fatal("Failed to consume stocks data", zap.Error(extErr))
	}

	server := &http.Server{
		Addr:    config.Get().ListenURL,
		Handler: handlers.NewRouter(),
	}

	logger.Get().Info("Listening...", zap.String("listen_url", config.Get().ListenURL))
	err = server.ListenAndServe()
	if err != nil {
		logger.Get().Error("Failed to initialize HTTP server", zap.Error(err))
		os.Exit(1)
	}
}
