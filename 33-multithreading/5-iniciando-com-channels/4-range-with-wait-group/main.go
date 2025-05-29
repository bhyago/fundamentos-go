package main

import "sync"

func main() {
	ch := make(chan int) // Create a channel to communicate between goroutines'
	wg := sync.WaitGroup{} // Create a WaitGroup to wait for goroutines to finish
	wg.Add(10) // Add one goroutine to the WaitGroup

	go publish(ch) // Start the publisher goroutine
	go reader(ch, &wg) // Start the reader goroutine
	wg.Wait()
}

func reader(ch chan int , wg *sync.WaitGroup) {
	for x := range ch {
		println(x)
		wg.Done() // Mark the goroutine as done
	}
}
// The reader function reads from the channel until it is closed.

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all values
}