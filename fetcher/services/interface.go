package services

import (
	"github.com/eugeneradionov/stocks/fetcher/services/rabbitmq"
	"github.com/eugeneradionov/stocks/fetcher/services/stocks"
)

type Service interface {
	Stocks() stocks.Service
	RabbitMQ() rabbitmq.Service
}
