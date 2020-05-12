package symbols

import (
	"context"
	"encoding/json"

	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (srv service) Consume() exterrors.ExtError {
	msgs, err := srv.symbolsCh.Consume(
		srv.rabbitSrv.GetSymbolsQueue().Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "rabbitMQ consume symbols"))
	}

	go func() {
		for delivery := range msgs {
			extErr := srv.processSymbols(delivery)
			if extErr != nil {
				logger.Get().Error("failed to process symbols", zap.Error(extErr))
			}
		}
	}()

	return nil
}

func (srv service) processSymbols(d amqp.Delivery) exterrors.ExtError {
	var symbols []models.Symbol

	err := json.Unmarshal(d.Body, &symbols)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "unmarshal symbols"))
	}

	symMap := make(map[string]struct{}, len(symbols))
	uniqueSymbols := make([]models.Symbol, 0, len(symbols))

	var ok bool

	for i := range symbols {
		if _, ok = symMap[symbols[i].Symbol]; ok {
			continue
		}

		uniqueSymbols = append(uniqueSymbols, symbols[i])
		symMap[symbols[i].Symbol] = struct{}{}
	}

	tx, err := postgres.GetDB().NewTXContext(context.Background())
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "start db transaction"))
	}
	defer tx.RollbackTx("process symbols")

	err = repo.Get().Symbols().
		WithTx(tx).
		BulkUpsert(uniqueSymbols)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "symbols bulk upsert"))
	}

	err = d.Ack(false)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "ack symbols message"))
	}

	err = tx.Commit()
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "commit db transaction"))
	}

	return nil
}
