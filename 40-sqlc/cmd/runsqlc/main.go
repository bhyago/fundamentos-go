package main

import (
	"context"
	"database/sql"
	"sqlcgo/internal/db"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3303)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:   uuid.New().String(),
	// 	Name: "New Category",
	// 	Description: sql.NullString{
	// 		String: "This is a new category", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println("Category ID:", category.ID)
	// 	println("Category Name:", category.Name)
	// 	if category.Description.Valid {
	// 		println("Category Description:", category.Description.String)
	// 	} else {
	// 		println("Category Description: NULL")
	// 	}
	// }

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "833bcae2-fc2b-4d4c-ba96-a6f43e87ffe2",
		Name:        "Updated Category",
		Description: sql.NullString{String: "This is an updated category", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println("Category ID:", category.ID)
		println("Category Name:", category.Name)
		if category.Description.Valid {
			println("Category Description:", category.Description.String)
		} else {
			println("Category Description: NULL")
		}
	}

	err = queries.DeleteCategory(ctx, "833bcae2-fc2b-4d4c-ba96-a6f43e87ffe2")
	if err != nil {
		panic(err)
	}

	category, err := queries.GetCategory(ctx, "833bcae2-fc2b-4d4c-ba96-a6f43e87ffe2")
	if err != nil {
		if err == sql.ErrNoRows {
			println("Category not found")
		} else {
			panic(err)
		}
	} else {
		println("Category found:", category.ID)
		println("Category Name:", category.Name)
		if category.Description.Valid {
			println("Category Description:", category.Description.String)
		} else {
			println("Category Description: NULL")
		}
	}
}
