package main

import "fmt"

func recebe(nome string, hello chan<- string) {
	hello <- nome
}

func ler(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	// hello <- "Hello World!" // This will cause a deadlock because the channel is unbuffered and there is no receiver
	// go recebe("Hello World!", hello) // This will also cause a deadlock because the channel is unbuffered and there is no receiver
	go recebe("Hello World!", hello) // Start the goroutine to send data to the channel
	ler(hello) // Start the goroutine to receive data from the channel

}