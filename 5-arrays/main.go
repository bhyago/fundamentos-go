package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 12
	meuArray[1] = 21
	meuArray[2] = 32

	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice %d é o valor é %d\n", i, v)
	}
}
