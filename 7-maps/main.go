package main

import "fmt"

func main() {
	salarios := map[string]int{"hyago": 100, "rayane": 1120, "solange": 2323}
	fmt.Println(salarios["hyago"])
	delete(salarios, "rayane")
	salarios["Rayane"] = 500

	fmt.Println(salarios["Rayane"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salrio de %s é %d\n", nome, salario)
	}
}
