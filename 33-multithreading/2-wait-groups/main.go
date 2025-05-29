package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done() // Decrement the wait group counter
	}
}

func main() {
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(25)

	go task("A", &waitgroup)
	go task("B", &waitgroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Anonymous function is running\n", i)
			time.Sleep(1 * time.Second)
			waitgroup.Done() // Decrement the wait group counter
		}
	}()

	waitgroup.Wait()
}
