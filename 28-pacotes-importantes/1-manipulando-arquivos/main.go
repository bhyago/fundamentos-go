package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Hello, World!")
	tamanho, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criando com sucesso! Tamanho: %d bytes", tamanho)
	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
