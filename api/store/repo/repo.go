package repo

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/store/repo/postgres/symbols"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	symbolsRepo symbols.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			symbolsRepo: symbols.NewSymbolsDAO(),
		}
	})

	return err
}

func (r postgresRepo) Symbols() symbols.DAO {
	return r.symbolsRepo
}
