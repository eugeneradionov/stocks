package stocks

import (
	"encoding/json"
	"time"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/models"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func (srv service) ProcessCandlesRPC(ch *amqp.Channel, msg amqp.Delivery) exterrors.ExtError {
	var req models.CandleRPCRequest

	err := json.Unmarshal(msg.Body, &req)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "unmarshal candle rpc request"))
	}

	from := time.Unix(req.From, 0)
	to := time.Unix(req.To, 0)

	candles, extErr := srv.GetCandles(req.Symbol, req.Resolution, from, to)
	if extErr != nil {
		return extErr
	}

	respBody, err := json.Marshal(&candles)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "marshal candles rpc response body"))
	}

	err = ch.Publish(
		"",
		msg.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			CorrelationId:   msg.CorrelationId,
			Body:            respBody,
		},
	)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "publish candles rpc response"))
	}

	err = msg.Ack(false)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "ack candles rpc request"))
	}

	return nil
}
