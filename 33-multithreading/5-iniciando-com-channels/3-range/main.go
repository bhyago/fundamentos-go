package main

func main() {
	ch := make(chan int) // Create a channel to communicate between goroutines'
	go publish(ch) // Start the publisher goroutine
	reader(ch) // Start the reader goroutine
}

func reader(ch chan int) {
	for x := range ch {
		println(x)
	}
}
// The reader function reads from the channel until it is closed.

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all values
}