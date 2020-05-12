package rabbitmq

import "github.com/streadway/amqp"

func (srv service) GetExchangesQueue() *amqp.Queue {
	return srv.exchangesQueue
}

func (srv service) DeclareExchangesQueue() (amqp.Queue, error) {
	return srv.rabbitCli.GetCh().QueueDeclare(
		"stocks.exchanges",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (srv service) GetSymbolsQueue() *amqp.Queue {
	return srv.symbolsQueue
}

func (srv service) DeclareSymbolsQueue() (amqp.Queue, error) {
	return srv.rabbitCli.GetCh().QueueDeclare(
		"stocks.symbols",
		false,
		false,
		false,
		false,
		nil,
	)
}
