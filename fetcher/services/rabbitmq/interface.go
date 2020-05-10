package rabbitmq

import "github.com/streadway/amqp"

type Service interface {
	GetExchangesQueue() *amqp.Queue
	GetSymbolsQueue() *amqp.Queue

	DeclareExchangesQueue() (amqp.Queue, error)
	DeclareSymbolsQueue() (amqp.Queue, error)

	Publish(exchange, queueName string, mandatory, immediate bool, body []byte) error
}
