package main

import "fmt"

func main() {
	slice := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	fmt.Printf("len=%d cap%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("len=%d cap%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	fmt.Printf("len=%d cap%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	slice = append(slice, 12)
	fmt.Printf("len=%d cap%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])
}
