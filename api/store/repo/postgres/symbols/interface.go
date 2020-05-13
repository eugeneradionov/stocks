package symbols

import (
	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type DAO interface {
	// WithTx returns new DAO with injected tx
	WithTx(tx *postgres.DBQuery) DAO

	// Get returns list of symbols with pagination
	Get(limit, offset int) (symbols []models.Symbol, err error)

	// GetByName returns symbol by name if any
	GetByName(name string) (symbol models.Symbol, err error)

	Insert(symbol models.Symbol) (_ models.Symbol, err error)
	BulkInsert(symbols []models.Symbol) (err error)

	Upsert(symbol models.Symbol) (_ models.Symbol, err error)
	BulkUpsert(symbols []models.Symbol) (err error)
}
