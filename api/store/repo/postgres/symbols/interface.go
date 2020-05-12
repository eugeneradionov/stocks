package symbols

import (
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	GetByName(name string) (symbol models.Symbol, err error)

	Insert(symbol models.Symbol) (_ models.Symbol, err error)
	BulkInsert(symbols []models.Symbol) (err error)

	Upsert(symbol models.Symbol) (_ models.Symbol, err error)
	BulkUpsert(symbols []models.Symbol) (err error)
}
