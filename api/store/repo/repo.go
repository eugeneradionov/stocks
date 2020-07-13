package repo

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/store/repo/postgres/candles"
	logins_session "github.com/eugeneradionov/stocks/api/store/repo/postgres/logins-session"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/users"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	loginSessionDAO logins_session.DAO
	usersDAO        users.DAO

	symbolsDAO symbols.DAO
	candlesDAO candles.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			loginSessionDAO: logins_session.NewLoginSessionDAO(),
			usersDAO:        users.NewUsersDAO(),
			symbolsDAO:      symbols.NewSymbolsDAO(),
			candlesDAO:      candles.NewCandlesDAO(),
		}
	})

	return err
}

func (r postgresRepo) Symbols() symbols.DAO {
	return r.symbolsDAO
}

func (r postgresRepo) Candles() candles.DAO {
	return r.candlesDAO
}

func (r postgresRepo) LoginSession() logins_session.DAO {
	return r.loginSessionDAO
}

func (r postgresRepo) Users() users.DAO {
	return r.usersDAO
}
