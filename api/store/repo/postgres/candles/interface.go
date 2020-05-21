package candles

import (
	"time"

	"github.com/eugeneradionov/stocks/api/models"
)

type DAO interface {
	Get(symbol string, resolution models.CandleResolution, from, to time.Time) (models.Candle, error)

	Insert(candle models.Candle) (_ models.Candle, err error)
}
