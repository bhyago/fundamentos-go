package main

import (
	"fmt"

	"github.com/devfullcycle/fcutils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println("Event Dispatcher created:", ed)

}
