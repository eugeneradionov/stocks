package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	cli *redis.Client
}

type Cli interface {
	// Get returns value from redis by `key`.
	Get(key string) ([]byte, error)
	// GetInt returns int values from redis by `key`.
	GetInt(key string) (int, error)
	// GetString returns string values from redis by `key`.
	GetString(key string) string

	// Set sets value to redis by `key` with `ttl`.
	Set(key string, value []byte, ttl time.Duration) error
	// Set sets int value to redis by `key` with `ttl`.
	SetInt(key string, value int, ttl time.Duration) error

	// Del deletes `keys` from redis.
	Del(keys ...string) error

	// TTL returns TTL by the `key`.
	TTL(key string) (time.Duration, error)
	// Expire adds expiration(`exp`) time for `key`.
	Expire(key string, exp time.Duration) error
	// Exists checks if a `key` exists.
	Exists(key string) (bool, error)

	// Ping pings the redis.
	Ping(ctx context.Context) error
}

func (r *Client) Get(key string) ([]byte, error) {
	res, err := r.cli.Get(key).Bytes()

	if err == redis.Nil {
		// it means that key does not exist in redis
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *Client) GetInt(key string) (int, error) {
	res, err := r.cli.Get(key).Int()
	if err == redis.Nil {
		// it means that key does not exist in redis
		return 0, nil
	}

	return res, err
}

func (r *Client) GetString(key string) string {
	return r.cli.Get(key).String()
}

func (r *Client) Set(key string, value []byte, ttl time.Duration) error {
	return r.cli.Set(key, value, ttl).Err()
}

func (r *Client) SetInt(key string, value int, ttl time.Duration) error {
	return r.cli.Set(key, value, ttl).Err()
}

func (r *Client) Del(keys ...string) error {
	return r.cli.Del(keys...).Err()
}

func (r *Client) TTL(key string) (time.Duration, error) {
	return r.cli.TTL(key).Result()
}

func (r *Client) Expire(key string, exp time.Duration) error {
	return r.cli.Expire(key, exp).Err()
}

func (r *Client) Exists(key string) (bool, error) {
	res, err := r.cli.Exists(key).Result()
	if err != nil {
		return false, err
	}

	if res == 1 {
		return true, nil
	}

	return false, nil
}

func (r *Client) Ping(ctx context.Context) error {
	return r.cli.WithContext(ctx).Ping().Err()
}
