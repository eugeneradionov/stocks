package repo

import "github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"

type Repo interface {
	Symbols() symbols.DAO
}
