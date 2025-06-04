package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	useCase := NewUseCase(db)

	product, err := useCase.GetProduct(1)
	println("Product ID:", product.ID, "Name:", product.Name)
}
