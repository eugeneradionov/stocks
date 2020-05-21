package rabbitmq

import "github.com/streadway/amqp"

type Service interface {
	NewCh() (*amqp.Channel, error)
	CloseCh(ch *amqp.Channel)

	GetExchangesQueue() *amqp.Queue
	GetSymbolsQueue() *amqp.Queue

	DeclareExchangesQueue() (amqp.Queue, error)
	DeclareSymbolsQueue() (amqp.Queue, error)

	Publish(exchange, queueName string, mandatory, immediate bool, body []byte) error
	RPCCall(routingKey string, reqBody []byte) ([]byte, error)
}
