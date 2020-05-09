package stocks

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
)

func (srv service) GetExchanges() (exchanges []models.Exchange, extErr exterrors.ExtError) {
	return srv.adapter.GetExchanges()
}
