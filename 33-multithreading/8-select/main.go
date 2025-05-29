package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}

	}()
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msd := Message{i, "Hello from Kafka"}
			c2 <- msd
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case ms1 := <-c1:
			fmt.Printf("Received from c1: %d %s\n", ms1.id, ms1.Msg)
		case ms2 := <-c2:
			fmt.Printf("Received from c2: %d %s\n", ms2.id, ms2.Msg)
		case <-time.After(time.Second * 3):
			println("Timeout")
			// default:
			// println("No messages received")
		}

	}
}
