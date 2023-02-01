package main

import (
	"fmt"

	"github.com/giovane-aG/goexpert/goevents/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	msgChannel := make(chan amqp.Delivery)
	defer ch.Close()

	go rabbitmq.Consume(ch, msgChannel, "orders")

	for msg := range msgChannel {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
