package finnhub

import (
	"encoding/json"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/pkg/errors"
)

const exchangesPath = "stock/exchange"

func (f Finnhub) GetExchanges() (exchanges []models.Exchange, extErr exterrors.ExtError) {
	data, extErr := f.sendRequest(exchangesPath, nil)
	if extErr != nil {
		return nil, extErr
	}

	err := json.Unmarshal(data, &exchanges)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "finnhub unmarshal exchanges json body"),
		)
	}

	return exchanges, nil
}
