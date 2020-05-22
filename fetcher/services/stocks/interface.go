package stocks

import (
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/streadway/amqp"
)

type Service interface {
	// GetExchanges returns available exchanges list.
	GetExchanges() (exchanges []models.Exchange, extErr exterrors.ExtError)

	// GetSymbols publishes available symbols for exchange to rabbit
	GetSymbols(exchangeCode string) (extErr exterrors.ExtError)

	// GetCandles returns candles for symbol with resolution for period
	GetCandles(symbol string, resolution models.CandleResolution,
		from, to time.Time) (candles models.Candle, extErr exterrors.ExtError)

	// ProcessCandlesRPC handles candles RPC and writes back the response
	ProcessCandlesRPC(ch *amqp.Channel, msg amqp.Delivery) exterrors.ExtError

	// RefreshSymbols spawns goroutine and gets symbols for exchange code with specified timeout.
	RefreshSymbols(cfg config.Symbols, exchangeCode string)
}
