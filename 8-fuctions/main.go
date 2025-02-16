package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(1, 2))

	valor, err := sum4(1, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

func sum4(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("a soma é maior que 50")
	}

	return a + b, nil
}
