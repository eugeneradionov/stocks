package services

import (
	"github.com/eugeneradionov/stocks/api/services/auth"
	"github.com/eugeneradionov/stocks/api/services/candles"
	"github.com/eugeneradionov/stocks/api/services/password"
	"github.com/eugeneradionov/stocks/api/services/stocks"
	"github.com/eugeneradionov/stocks/api/services/symbols"
	"github.com/eugeneradionov/stocks/api/services/token"
)

type Service interface {
	Stocks() stocks.Service

	Symbols() symbols.Service
	Candles() candles.Service

	Auth() auth.Service
	Token() token.Service
	Password() password.Service
}
