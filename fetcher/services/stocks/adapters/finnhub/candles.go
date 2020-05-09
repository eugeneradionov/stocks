package finnhub

import (
	"encoding/json"
	"strconv"
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/pkg/errors"
)

const candlesPath = "stock/candle"

func (f Finnhub) GetCandles(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (candles models.Candle, extErr exterrors.ExtError) {
	data, extErr := f.sendCandlesRequest(symbol, resolution, from, to)
	if extErr != nil {
		return candles, extErr
	}

	err := json.Unmarshal(data, &candles)
	if err != nil {
		return candles, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "finnhub unmarshal candles json body"),
		)
	}

	return candles, nil
}

func (f Finnhub) sendCandlesRequest(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) ([]byte, exterrors.ExtError) {
	params := map[string]string{
		"symbol":     symbol,
		"resolution": string(resolution),
		"from":       strconv.FormatInt(from.UTC().Unix(), 10),
		"to":         strconv.FormatInt(to.UTC().Unix(), 10),
	}

	return f.sendRequest(candlesPath, params)
}
