package main

func main() {
	ch := make(chan string, 2) // Create a buffered channel with a capacity of 2
	ch <- "Hello"
	ch <- "World"

	print(<-ch) // Read from the channel
	print(<-ch) // Read from the channel
}
