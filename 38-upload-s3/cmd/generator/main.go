package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		f, error := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if error != nil {
			panic(error)
		}
		f.Close()
		f.WriteString("Hello, World!\n")
		i++
	}
}
