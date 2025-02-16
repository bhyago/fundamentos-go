package main

import (
	"fmt"

	"github.com/bhyago/fundamentos-go/21-pacotes-e-modulos/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado:", s)
	fmt.Println(uuid.New())
}
