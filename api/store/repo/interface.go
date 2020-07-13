package repo

import (
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/candles"
	logins_session "github.com/eugeneradionov/stocks/api/store/repo/postgres/logins-session"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/users"
)

type Repo interface {
	LoginSession() logins_session.DAO
	Users() users.DAO
	Symbols() symbols.DAO
	Candles() candles.DAO
}
