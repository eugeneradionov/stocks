package services

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/services/auth"
	"github.com/eugeneradionov/stocks/api/services/candles"
	"github.com/eugeneradionov/stocks/api/services/password"
	"github.com/eugeneradionov/stocks/api/services/rabbitmq"
	"github.com/eugeneradionov/stocks/api/services/stocks"
	"github.com/eugeneradionov/stocks/api/services/symbols"
	"github.com/eugeneradionov/stocks/api/services/token"
	"github.com/eugeneradionov/stocks/api/store/repo"
	"github.com/eugeneradionov/stocks/api/store/repo/redis"
	rabbit "github.com/eugeneradionov/stocks/api/transport/rabbitmq"
)

var (
	srvRepo serviceRepo
	once    = &sync.Once{}
)

type serviceRepo struct {
	rabbitSrv   rabbitmq.Service
	symbolsSrv  symbols.Service
	candlesSrv  candles.Service
	stocksSrv   stocks.Service
	tokenSrv    token.Service
	authSrv     auth.Service
	passwordSrv password.Service
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

func (srv serviceRepo) Candles() candles.Service {
	return srv.candlesSrv
}

func (srv serviceRepo) Token() token.Service {
	return srv.tokenSrv
}

func (srv serviceRepo) Auth() auth.Service {
	return srv.authSrv
}

func (srv serviceRepo) Password() password.Service {
	return srv.passwordSrv
}

func Load(cfg *config.Config, redisCli redis.Cli, rabbitCli *rabbit.Rabbit) (err error) {
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

		candlesSrv, e := candles.New(rabbitSrv)
		if e != nil {
			err = e
			return
		}

		tokenSrv := token.New(redisCli, cfg.Token)
		passwordSrv := password.New()
		authSrv := auth.New(tokenSrv, passwordSrv, repo.Get().Users(), repo.Get().LoginSession())

		srvRepo = serviceRepo{
			rabbitSrv:   rabbitSrv,
			symbolsSrv:  symbolsSrv,
			candlesSrv:  candlesSrv,
			stocksSrv:   stocksSrv,
			tokenSrv:    tokenSrv,
			authSrv:     authSrv,
			passwordSrv: passwordSrv,
		}
	})

	return err
}
