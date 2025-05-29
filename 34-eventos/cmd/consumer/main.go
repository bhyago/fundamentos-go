package main

import (
	"eventos/pkg/rabbitmq"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp091.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Printf("Mensagem recebida: %s\n", msg.Body)
		msg.Ack(false)
	}
}
