package main

import "time"

func worker(workerID int, data chan int) {
	for item := range data {
		// Simulate work
		println("Worker", workerID, "processing item", item)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 10
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}
	// Simulate sending data to the workers
	// In a real-world scenario, this could be a stream of data
	// from a database, a message queue, etc.
	// Here, we are just sending numbers from 0 to 99
	// to simulate work
	// and then closing the channel to signal the workers to stop
	// processing
	// the data

	for i := 0; i < 100; i++ {
		data <- i
	}
	close(data)
}