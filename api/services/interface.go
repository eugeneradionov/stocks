package services

import "github.com/eugeneradionov/stocks/api/services/stocks"

type Service interface {
	Stocks() stocks.Service
}
