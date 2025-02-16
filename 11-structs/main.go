package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	hyago := Cliente{
		Nome:  "Hyago",
		Idade: 20,
		Ativo: true,
	}

	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", hyago.Nome, hyago.Idade, hyago.Ativo)
	hyago.Ativo = false
	fmt.Println(hyago.Ativo)
}
