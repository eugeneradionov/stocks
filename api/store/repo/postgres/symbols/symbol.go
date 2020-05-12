package symbols

import (
	"context"

	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type symbolsDAO struct {
	q *postgres.DBQuery
}

func NewSymbolsDAO() DAO {
	return &symbolsDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao symbolsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &symbolsDAO{q: tx}
}

func (dao symbolsDAO) GetByName(name string) (symbol models.Symbol, err error) {
	err = dao.q.Model(&symbol).
		Where("symbol = ?", name).
		First()

	return symbol, err
}

func (dao symbolsDAO) Insert(symbol models.Symbol) (_ models.Symbol, err error) {
	_, err = dao.q.Model(&symbol).
		Returning("*").
		Insert()

	return symbol, err
}

func (dao symbolsDAO) BulkInsert(symbols []models.Symbol) (err error) {
	if len(symbols) < 1 {
		return nil
	}

	_, err = dao.q.Model(&symbols).Insert()

	return err
}

func (dao symbolsDAO) Upsert(symbol models.Symbol) (_ models.Symbol, err error) {
	_, err = dao.q.Model(&symbol).
		OnConflict("(symbol) DO UPDATE").
		Returning("*").
		Insert()

	return symbol, err
}

func (dao symbolsDAO) BulkUpsert(symbols []models.Symbol) (err error) {
	_, err = dao.q.Model(&symbols).
		OnConflict("(symbol) DO UPDATE").
		Insert()

	return err
}
