package redis

import (
	"sync"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/go-redis/redis"
)

var (
	client *Client
	once   = &sync.Once{}
)

func Load(cfg config.Redis) (err error) {
	once.Do(func() {
		cli := redis.NewClient(
			&redis.Options{
				Addr:     cfg.Address,
				Password: cfg.Password,
				PoolSize: cfg.PoolSize,
			})

		err = cli.Ping().Err()
		if err != nil {
			return
		}

		client = &Client{cli: cli}
	})

	return err
}

func GetRedis() Cli {
	return client
}
