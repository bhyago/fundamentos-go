package main

import (
	"fmt"

	"github.com/bhyago/fundamentos-go/31-packaging/1-intro-go-mod/math"
)

func main() {
	// m := math.Math{A: 1, B: 2}
	m := math.Math{}
	fmt.Println(m.Add())
	fmt.Println("Hello, world!")
}
