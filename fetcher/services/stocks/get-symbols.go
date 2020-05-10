package stocks

import (
	"encoding/json"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (srv service) GetSymbols(exchangeCode string) (extErr exterrors.ExtError) {
	symbols, extErr := srv.adapter.GetSymbols(exchangeCode)
	if extErr != nil {
		return extErr
	}

	logger.Get().Info("publishing symbols to rabbit",
		zap.String("exchange_code", exchangeCode), zap.Int("symbols_length", len(symbols)))

	extErr = srv.publishSymbols(symbols)
	if extErr != nil {
		return extErr
	}

	return nil
}

func (srv service) publishSymbols(symbols []models.Symbol) exterrors.ExtError {
	data, err := json.Marshal(symbols)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "marshal symbols json body"))
	}

	err = srv.rabbitSrv.Publish(
		"",
		srv.rabbitSrv.GetSymbolsQueue().Name,
		false,
		false,
		data,
	)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "publish symbols to rabbitMQ"))
	}

	return nil
}
