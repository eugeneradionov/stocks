package candles

import (
	"context"
	"time"

	"github.com/eugeneradionov/stocks/api/models"
	"github.com/eugeneradionov/stocks/api/store/repo/postgres"
)

type candlesDAO struct {
	q *postgres.DBQuery
}

func NewCandlesDAO() DAO {
	return &candlesDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao candlesDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &candlesDAO{q: tx}
}

func (dao candlesDAO) Get(
	symbol string,
	resolution models.CandleResolution,
	from, to time.Time,
) (candle models.Candle, err error) {
	err = dao.q.Model(&candle).
		Where("symbol = ?", symbol).
		Where("resolution = ?", resolution).
		Where(`"from" = ?`, from).
		Where(`"to" = ?`, to).
		Select()

	return candle, err
}

func (dao candlesDAO) Insert(candle models.Candle) (_ models.Candle, err error) {
	_, err = dao.q.Model(&candle).
		Returning("*").
		Insert()

	return candle, err
}
