package main

import (
	"log"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/services"
	"github.com/eugeneradionov/stocks/api/store/repo"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/eugeneradionov/stocks/api/transport/rabbitmq"
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
		panic(extErr)
	}

	// TODO: remove
	forever := make(chan struct{})
	<-forever

	// symbol := models.Symbol{
	// 	Description:   "test description",
	// 	DisplaySymbol: "TST",
	// 	Symbol:        "TST",
	// }
	// res, err := repo.Get().Symbols().Insert(symbol)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// res, err = repo.Get().Symbols().GetByName("TST")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(res) // TODO: REMOVE
}
