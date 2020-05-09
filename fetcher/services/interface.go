package services

import "github.com/eugeneradionov/stocks/fetcher/services/stocks"

type Service interface {
	Stocks() stocks.Service
}
