package services

import (
	"github.com/eugeneradionov/stocks/api/services/stocks"
	"github.com/eugeneradionov/stocks/api/services/symbols"
)

type Service interface {
	Stocks() stocks.Service
	Symbols() symbols.Service
}
