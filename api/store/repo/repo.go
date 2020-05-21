package repo

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/store/repo/postgres/candles"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	symbolsDAO symbols.DAO
	candlesDAO candles.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			symbolsDAO: symbols.NewSymbolsDAO(),
			candlesDAO: candles.NewCandlesDAO(),
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
