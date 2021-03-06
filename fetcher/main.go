package main

import (
	"log"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"github.com/eugeneradionov/stocks/fetcher/services"
	"github.com/eugeneradionov/stocks/fetcher/transport/rabbitmq"
	"go.uber.org/zap"
)

const defaultConfigPath = "config.json"

func main() {
	err := config.Load(defaultConfigPath)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err.Error())
	}

	err = logger.Load(config.Get())
	if err != nil {
		log.Fatalf("Failed to laod logger: %s", err.Error())
	}

	err = rabbitmq.Load(config.Get().RabbitMQ)
	if err != nil {
		logger.Get().Fatal("Failed to connect to rabbit", zap.Error(err))
	}

	err = services.Load(config.Get(), rabbitmq.Get())
	if err != nil {
		logger.Get().Fatal("Failed to load services", zap.Error(err))
	}

	// setup delayed job for symbols refreshing for US exchange
	services.Get().Stocks().RefreshSymbols(config.Get().Stocks.Symbols, "US")

	go func() {
		e := services.Get().RabbitMQ().ListenCandlesRPC(services.Get().Stocks().ProcessCandlesRPC)
		if e != nil {
			logger.Get().Fatal("failed to start candles RPC listener", zap.Error(e))
		}
	}()

	logger.Get().Info("Services started successfully")

	forever := make(chan struct{})
	<-forever
}
