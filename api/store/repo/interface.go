package repo

import (
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/candles"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"
)

type Repo interface {
	Symbols() symbols.DAO
	Candles() candles.DAO
}
