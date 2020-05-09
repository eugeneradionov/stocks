package services

import (
	"sync"

	"github.com/eugeneradionov/stocks/fetcher/config"
	"github.com/eugeneradionov/stocks/fetcher/services/stocks"
	"github.com/eugeneradionov/stocks/fetcher/transport/http"
)

var (
	srv  serviceRepo
	once = &sync.Once{}
)

type serviceRepo struct {
	stocks stocks.Service
}

func Get() Service {
	return srv
}

func (srv serviceRepo) Stocks() stocks.Service {
	return srv.stocks
}

func Load(cfg *config.Config) (err error) {
	once.Do(func() {
		stocksSrv, e := stocks.New(cfg.Stocks, http.NewClient())
		if e != nil {
			err = e
		}

		srv = serviceRepo{
			stocks: stocksSrv,
		}
	})

	return err
}
