package rabbitmq

import (
	"sync"

	"github.com/eugeneradionov/stocks/fetcher/transport/rabbitmq"
	"github.com/streadway/amqp"
)

type service struct {
	rabbitCli      *rabbitmq.Rabbit
	exchangesQueue *amqp.Queue
	symbolsQueue   *amqp.Queue
}

var (
	srv  service
	once = &sync.Once{}
)

func New(rabbitCli *rabbitmq.Rabbit) (_ Service, err error) {
	once.Do(func() {
		var (
			exchangesQueue amqp.Queue
			symbolsQueue   amqp.Queue
		)

		srv = service{
			rabbitCli: rabbitCli,
		}

		exchangesQueue, err = srv.DeclareExchangesQueue()
		if err != nil {
			return
		}

		srv.exchangesQueue = &exchangesQueue

		symbolsQueue, err = srv.DeclareSymbolsQueue()
		if err != nil {
			return
		}

		srv.symbolsQueue = &symbolsQueue
	})

	return srv, err
}
