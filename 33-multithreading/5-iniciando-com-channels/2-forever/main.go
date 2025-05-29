package main

func main() {
	forever := make(chan bool) // Create a channel to block the main thread

	go func  () {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // Send a value to the channel to unblock the main thread
	}()

	<-forever // Block the main thread until a value is received from the channel
}