package finnhub

import (
	"encoding/json"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/pkg/errors"
)

const symbolsPath = "stock/symbol"

func (f Finnhub) GetSymbols(exchangeCode string) (symbols []models.Symbol, extErr exterrors.ExtError) {
	data, extErr := f.sendSymbolsRequest(exchangeCode)
	if extErr != nil {
		return nil, extErr
	}

	err := json.Unmarshal(data, &symbols)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "finnhub unmarshal symbols json body"),
		)
	}

	return symbols, nil
}

func (f Finnhub) sendSymbolsRequest(exchangeCode string) ([]byte, exterrors.ExtError) {
	return f.sendRequest(symbolsPath, map[string]string{"exchange": exchangeCode})
}
