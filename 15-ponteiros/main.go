package main

import "fmt"

func main() {

	// Memoria > Endereço > valor
	a := 10

	//valor
	println(a)

	//endereço da memoria
	println(&a)

	// * vai armazenar o endereço de memoria
	// ponteiro é o endereçamento na memeria
	// variavel > ponteiro que tem um endereço na memeria > que tem um valor
	var ponteiro *int = &a

	//valor do endereço (poteiro) atualizado
	*ponteiro = 20

	println(a)
	b := &a
	*b = 1

	// pega o valor do endetenõ da memeria
	fmt.Println(*b)

	fmt.Println(a)

}
