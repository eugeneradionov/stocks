package candles

import (
	"encoding/json"
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/go-pg/pg/v9"
	"github.com/pkg/errors"
)

const candlesRoutingKey = "candles_rpc_queue"

func (srv service) Get(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (models.Candle, exterrors.ExtError) {
	_, err := srv.symbolsDAO.GetByName(symbol)
	if err != nil {
		if err == pg.ErrNoRows {
			return models.Candle{}, exterrors.NewNotFoundError(
				errors.New("symbol not found"), "symbol_name",
			)
		}

		return models.Candle{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get symbol by name"))
	}

	var extErr exterrors.ExtError

	candles, err := srv.candlesDAO.Get(symbol, resolution, from, to)
	if err != nil {
		if err == pg.ErrNoRows {
			candles, extErr = srv.getAndInsert(symbol, resolution, from, to)
			if extErr != nil {
				return models.Candle{}, extErr
			}

			return candles, nil
		}

		return models.Candle{}, exterrors.NewInternalServerErrorError(
			errors.Wrapf(err, "get candles; symbol: '%s', resolution: '%s', from: '%s', to: '%s'",
				symbol, resolution, from.String(), to.String(),
			),
		)
	}

	return candles, nil
}

func (srv service) getAndInsert(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (candles models.Candle, extErr exterrors.ExtError) {
	candles, extErr = srv.getWithPRC(symbol, resolution, from, to)
	if extErr != nil {
		return models.Candle{}, extErr
	}

	var err error

	candles, err = srv.candlesDAO.Insert(candles)
	if err != nil {
		return models.Candle{}, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "insert candles"),
		)
	}

	return candles, nil
}

func (srv service) getWithPRC(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (candles models.Candle, extErr exterrors.ExtError) {
	req := models.CandleRPCRequest{
		Symbol:     symbol,
		Resolution: resolution,
		From:       from.UTC().Unix(),
		To:         to.UTC().Unix(),
	}

	reqBody, err := json.Marshal(&req)
	if err != nil {
		return models.Candle{}, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "marshal candles RPC request"),
		)
	}

	respBody, err := srv.rabbitSrv.RPCCall(candlesRoutingKey, reqBody)
	if err != nil {
		return models.Candle{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get candles by RPC"))
	}

	var data models.CandleRabbit
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return models.Candle{}, exterrors.NewInternalServerErrorError(
			errors.Wrap(err, "unmarshal candles RPC response"),
		)
	}

	candles.Symbol = symbol
	candles.Resolution = resolution
	candles.From = from
	candles.To = to
	candles.Data = data.ToCandleData()

	return candles, nil
}
