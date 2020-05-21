package rabbitmq

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (srv service) ListenCandlesRPC(
	processFn func(ch *amqp.Channel, msg amqp.Delivery) exterrors.ExtError,
) (err error) {
	ch, err := srv.NewCh()
	if err != nil {
		return errors.Wrap(err, "create rabbitMQ channel")
	}
	defer srv.CloseCh(ch)

	q, err := ch.QueueDeclare(
		"candles_rpc_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "declare candles RPC queue")
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return errors.Wrap(err, "failed to configure QOS")
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "consume candles rpc request")
	}

	for msg := range msgs {
		extErr := processFn(ch, msg)
		if extErr != nil {
			logger.Get().Error("failed to process candles rpc request", zap.Error(extErr))
		}
	}

	return nil
}
