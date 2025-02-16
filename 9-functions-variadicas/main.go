package main

import "fmt"

func main() {
	fmt.Println(sum(1, 23, 45, 6, 56))
}

func sum(numeros ...int) int {
	total := 0
	for _, numer := range numeros {
		total += numer
	}

	return total
}
