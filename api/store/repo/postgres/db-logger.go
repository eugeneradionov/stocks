package postgres

import (
	"context"
	"time"

	"github.com/go-pg/pg/v9"
)

type queryStartedKey string

const queryStarted queryStartedKey = "query_started"

type DBLogger interface {
	Debug(args ...interface{})
	Debugf(fmt string, args ...interface{})
}

type dbLogger struct {
	logger DBLogger
}

// BeforeQuery logs before query execution.
func (d dbLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	ctx = context.WithValue(ctx, queryStarted, time.Now())
	d.logger.Debug(q.FormattedQuery())

	return ctx, nil
}

// AfterQuery logs after query execution.
func (d dbLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	execTime := time.Since(ctx.Value(queryStarted).(time.Time)).Round(time.Millisecond)
	if q.Err != nil {
		d.logger.Debugf("Query execution error: %s, exec time: %s", q.Err.Error(), execTime)
		return q.Err
	}

	if q.Result != nil {
		d.logger.Debugf("Rows affected: %d, Rows returned: %d, exec time: %s",
			q.Result.RowsAffected(), q.Result.RowsReturned(), execTime)
	}

	return nil
}
