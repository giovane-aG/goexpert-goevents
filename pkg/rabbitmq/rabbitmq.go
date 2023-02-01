package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"fila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, ctx context.Context) error {
	err := ch.PublishWithContext(
		ctx,
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
