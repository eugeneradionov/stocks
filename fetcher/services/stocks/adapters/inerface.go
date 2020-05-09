package adapters

import (
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
)

type Interface interface {
	// GetExchanges returns available exchanges list.
	GetExchanges() (exchanges []models.Exchange, extErr exterrors.ExtError)

	// GetSymbols returns available symbols for exchange
	GetSymbols(exchangeCode string) (symbols []models.Symbol, extErr exterrors.ExtError)

	// GetCandles returns candles for symbol with resolution for period
	GetCandles(symbol string, resolution models.CandleResolution,
		from, to time.Time) (candles models.Candle, extErr exterrors.ExtError)
}
