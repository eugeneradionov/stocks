package token

import (
	"time"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/eugeneradionov/stocks/api/store/repo"
	logins_session "github.com/eugeneradionov/stocks/api/store/repo/postgres/logins-session"
	"github.com/eugeneradionov/stocks/api/store/repo/redis"
)

type service struct {
	loginSessionDAO logins_session.DAO
	rds             redis.Cli

	expirationTimeSec int
	secretKey         string
}

func (srv service) MaxExpirationTime() int64 {
	return int64(srv.expirationTimeSec)
}

func New(rds redis.Cli, cfg config.Token) Service {
	return &service{
		loginSessionDAO:   repo.Get().LoginSession(),
		rds:               rds,
		expirationTimeSec: int(time.Duration(cfg.ExpirationTime).Seconds()),
		secretKey:         cfg.SecretKey,
	}
}
