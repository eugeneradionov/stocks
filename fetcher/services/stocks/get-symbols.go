package stocks

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
)

func (srv service) GetSymbols(exchangeCode string) (symbols []models.Symbol, extErr exterrors.ExtError) {
	return srv.adapter.GetSymbols(exchangeCode)
}
