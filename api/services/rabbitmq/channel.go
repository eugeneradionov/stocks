package rabbitmq

import "github.com/streadway/amqp"

func (srv service) NewCh() (*amqp.Channel, error) {
	return srv.rabbitCli.NewCh()
}
