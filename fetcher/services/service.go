package services

import (
	"sync"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/services/rabbitmq"
	"github.com/eugeneradionov/stocks/fetcher/services/stocks"
	"github.com/eugeneradionov/stocks/fetcher/transport/http"
	rabbit "github.com/eugeneradionov/stocks/fetcher/transport/rabbitmq"
)

var (
	srv  serviceRepo
	once = &sync.Once{}
)

type serviceRepo struct {
	stocks stocks.Service
	rabbit rabbitmq.Service
}

func Get() Service {
	return srv
}

func (srv serviceRepo) RabbitMQ() rabbitmq.Service {
	return srv.rabbit
}

func (srv serviceRepo) Stocks() stocks.Service {
	return srv.stocks
}

func Load(cfg *config.Config, rabbitCli *rabbit.Rabbit) (err error) {
	once.Do(func() {
		rabbitSrv, e := rabbitmq.New(rabbitCli)
		if e != nil {
			err = e
		}

		stocksSrv, e := stocks.New(cfg.Stocks, http.NewClient(), rabbitSrv)
		if e != nil {
			err = e
		}

		srv = serviceRepo{
			stocks: stocksSrv,
			rabbit: rabbitSrv,
		}
	})

	return err
}
