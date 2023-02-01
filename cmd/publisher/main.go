package main

import (
	"context"

	"github.com/giovane-aG/goexpert/goevents/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	ctx := context.Background()
	rabbitmq.Publish(ch, "Hello Rabbitmq", ctx)
}
