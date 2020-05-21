package rabbitmq

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

func (srv service) RPCCall(routingKey string, reqBody []byte) ([]byte, error) {
	ch, err := srv.NewCh()
	if err != nil {
		return nil, errors.Wrap(err, "new AMQP channel")
	}
	defer srv.CloseCh(ch)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "declare AMQP queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, errors.Wrap(err, "consume AMQP queue")
	}

	corrID := uuid.New().String()

	err = ch.Publish(
		"",
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:     "application/json",
			ContentEncoding: "utf-8",
			CorrelationId:   corrID,
			ReplyTo:         q.Name,
			Body:            reqBody,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "publish RPC message")
	}

	for msg := range msgs {
		if corrID == msg.CorrelationId {
			return msg.Body, nil
		}
	}

	return nil, errors.New("filed to get RPC response")
}
