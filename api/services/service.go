package services

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/services/rabbitmq"
	"github.com/eugeneradionov/stocks/api/services/stocks"
	"github.com/eugeneradionov/stocks/api/services/symbols"
	rabbit "github.com/eugeneradionov/stocks/api/transport/rabbitmq"
)

var (
	srvRepo serviceRepo
	once    = &sync.Once{}
)

type serviceRepo struct {
	rabbitSrv  rabbitmq.Service
	symbolsSrv symbols.Service
	stocksSrv  stocks.Service
}

func Get() Service {
	return srvRepo
}

func (srv serviceRepo) Stocks() stocks.Service {
	return srv.stocksSrv
}
func (srv serviceRepo) Symbols() symbols.Service {
	return srv.symbolsSrv
}

func Load(cfg *config.Config, rabbitCli *rabbit.Rabbit) (err error) {
	once.Do(func() {
		rabbitSrv, e := rabbitmq.New(rabbitCli)
		if e != nil {
			err = e
			return
		}

		symbolsSrv, e := symbols.New(rabbitSrv)
		if e != nil {
			err = e
			return
		}

		stocksSrv, e := stocks.New(symbolsSrv)
		if e != nil {
			err = e
			return
		}

		srvRepo = serviceRepo{
			rabbitSrv:  rabbitSrv,
			symbolsSrv: symbolsSrv,
			stocksSrv:  stocksSrv,
		}
	})

	return err
}
