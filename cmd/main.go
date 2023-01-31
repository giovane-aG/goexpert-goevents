package main

import (
	"fmt"

	"github.com/giovane-aG/goexpert/goevents/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	msgChannel := make(chan amqp.Delivery)

	err := rabbitmq.Consume(ch, msgChannel)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for msg := range msgChannel {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
