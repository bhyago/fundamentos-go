package main

//Thread 1
func main() {
	canal := make(chan string) // Vazio

	//Thread 2
	go func() {
		canal <- "Olá, mundo!" // Envia a mensagem para o canalW
		canal <- "Olá, mundo! 2" // Envia a mensagem para o canal
	}()

	//Thread 1
	mensagem := <-canal // Recebe a mensagem do canal
	println(mensagem)


} 