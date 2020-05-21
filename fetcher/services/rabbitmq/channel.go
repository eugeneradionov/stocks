package rabbitmq

import (
	"github.com/eugeneradionov/stocks/fetcher/logger"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func (srv service) NewCh() (*amqp.Channel, error) {
	return srv.rabbitCli.NewCh()
}

func (srv service) CloseCh(ch *amqp.Channel) {
	err := ch.Close()
	if err != nil {
		logger.Get().Error("failed to close AMQP channel", zap.Error(err))
	}
}
