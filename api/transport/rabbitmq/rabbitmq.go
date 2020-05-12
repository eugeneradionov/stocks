package rabbitmq

import (
	"fmt"
	"sync"

	"github.com/eugeneradionov/stocks/api/config"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

var (
	rabbit = &Rabbit{}
	once   = &sync.Once{}
)

func Load(cfg config.RabbitMQ) (err error) {
	once.Do(func() {
		rabbit.conn, err = amqp.Dial(rabbitURL(cfg))
		if err != nil {
			err = errors.Wrap(err, "dial to rabbitMQ")
		}

		rabbit.ch, err = rabbit.conn.Channel()
		if err != nil {
			err = errors.Wrap(err, "get rabbitMQ channel")
		}
	})

	return err
}

func Get() *Rabbit {
	return rabbit
}

// GetConn returns reserved AMQP connection.
func (r Rabbit) GetConn() *amqp.Connection {
	return r.conn
}

// GetCh returns reserved AMQP channel.
func (r Rabbit) GetCh() *amqp.Channel {
	return r.ch
}

// NewCh creates new AMQP channel with existing connection.
func (r Rabbit) NewCh() (*amqp.Channel, error) {
	return r.conn.Channel()
}

func rabbitURL(cfg config.RabbitMQ) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
}
