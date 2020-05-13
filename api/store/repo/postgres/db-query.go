package postgres

import (
	"context"
	"errors"

	"github.com/eugeneradionov/stocks/api/logger"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"go.uber.org/zap"
)

type DBModel interface {
	Model(model ...interface{}) *orm.Query
	Exec(query interface{}, params ...interface{}) (pg.Result, error)
}

type DBQuery struct {
	DBModel
	completed bool
}

// Rollback rollbacks query if it was transaction or returns error.
func (q *DBQuery) Rollback() error {
	switch t := q.DBModel.(type) { // nolint:gocritic
	case *pg.Tx:
		if !q.completed {
			return t.Rollback()
		}
		return nil
	}

	return errors.New("rollback failed: not in Tx")
}

// RollbackTx is the wrapper for graceful transaction rollback.
func (q *DBQuery) RollbackTx(info string) {
	err := q.Rollback()
	if err != nil {
		logger.Get().Error("failed to rollback transaction", zap.Error(err), zap.String("info", info))
	}
}

// Commit makes commit on transaction or do nothing.
func (q *DBQuery) Commit() error {
	switch t := q.DBModel.(type) { // nolint:gocritic
	case *pg.Tx:
		if !q.completed {
			q.completed = true

			return t.Commit()
		}
	}

	return nil
}

// NewTXContext returns DBQuery instance with new transaction
// DBQuery.Commit() must be called to run transaction.
func (p Postgres) NewTXContext(ctx context.Context) (*DBQuery, error) {
	tx, err := p.conn.WithContext(ctx).Begin()
	return &DBQuery{DBModel: tx}, err
}

// QueryContext returns DBQuery instance of current db pool.
func (p Postgres) QueryContext(ctx context.Context) *DBQuery {
	return &DBQuery{DBModel: p.conn.WithContext(ctx)}
}
