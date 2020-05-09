package main

import (
	"log"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"github.com/eugeneradionov/stocks/fetcher/services"
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

	err = services.Load(config.Get())
	if err != nil {
		logger.Get().Fatal("Failed to load services", zap.Error(err))
	}
}
