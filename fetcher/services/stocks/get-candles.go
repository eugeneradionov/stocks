package stocks

import (
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
)

func (srv service) GetCandles(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (candles models.Candle, extErr exterrors.ExtError) {
	return srv.adapter.GetCandles(symbol, resolution, from, to)
}
