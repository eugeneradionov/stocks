package postgres

import (
	"sync"
	"time"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/go-pg/pg/v9"
)

// Postgres pg connection.
type Postgres struct {
	conn *pg.DB
}

var (
	postgres *Postgres
	once     = &sync.Once{}
)

// Load - loads pg instance.
func Load(cfg config.Postgres, lgr DBLogger) error {
	once.Do(func() {
		db := pg.Connect(&pg.Options{
			Addr:         cfg.Host + ":" + cfg.Port,
			User:         cfg.User,
			Password:     cfg.Password,
			Database:     cfg.Database,
			PoolSize:     cfg.PoolSize,
			WriteTimeout: time.Duration(cfg.WriteTimeout),
			ReadTimeout:  time.Duration(cfg.ReadTimeout),
			MaxRetries:   cfg.MaxRetries,
		})

		db.AddQueryHook(dbLogger{
			logger: lgr,
		})

		postgres = &Postgres{
			conn: db,
		}
	})

	return postgres.Ping()
}

// GetDB returns pg DB instance.
func GetDB() *Postgres {
	return postgres
}

// Ping checks db connection.
func (p *Postgres) Ping() (err error) {
	var n int

	_, err = p.conn.QueryOne(pg.Scan(&n), "SELECT 1")

	return
}
