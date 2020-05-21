package candles

import (
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
)

type Service interface {
	Get(symbol string, resolution models.CandleResolution, from, to time.Time) (models.Candle, exterrors.ExtError)
}
