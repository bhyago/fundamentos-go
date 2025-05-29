package main

import "eventos/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	rabbitmq.Publish(ch, "Mensagem de teste", "amq.direct")
}
