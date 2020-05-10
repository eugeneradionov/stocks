package rabbitmq

import (
	"github.com/streadway/amqp"
)

func (srv service) Publish(
	exchange, routingKey string,
	mandatory, immediate bool,
	body []byte,
) error {
	return srv.rabbitCli.GetCh().Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			Body:            body,
		},
	)
}
