package symbols

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/services/rabbitmq"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type service struct {
	rabbitSrv rabbitmq.Service
	symbolsCh *amqp.Channel
}

var (
	srv  service
	once = &sync.Once{}
)

func New(rabbitSrv rabbitmq.Service) (_ Service, err error) {
	once.Do(func() {
		symbolsCh, e := rabbitSrv.NewCh()
		if e != nil {
			err = errors.Wrap(e, "rabbitMQ create symbols channel")
			return
		}

		srv = service{
			rabbitSrv: rabbitSrv,
			symbolsCh: symbolsCh,
		}
	})

	return srv, err
}
