package main

import (
	"context"
	"database/sql"
	"fmt"
	"sqlcgo/internal/db"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type CourseDB struct {
	dbConn  *sql.DB
	queries *db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		queries: db.New(dbConn),
	}
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	if err := fn(q); err != nil {
		if errb := tx.Rollback(); err != nil {
			return fmt.Errorf("tx rollback error: %v, original error: %w", errb, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return fmt.Errorf("create category error: %w", err)
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return fmt.Errorf("create course error: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction error: %w", err)
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3303)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	queries := db.New(dbConn)

	// coursesArgs := CourseParams{
	// 	ID:          "course-123",
	// 	Name:        "Go Programming",
	// 	Description: sql.NullString{String: "Learn Go from scratch", Valid: true},
	// 	Price:       29.99,
	// }

	// categoriesArgs := CategoryParams{
	// 	ID:          "category-123",
	// 	Name:        "Programming",
	// 	Description: sql.NullString{String: "All about programming", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)
	// err = courseDB.CreateCourseAndCategory(ctx, categoriesArgs, coursesArgs)
	// if err != nil {
	// 	fmt.Printf("Error creating course and category: %v\n", err)
	// 	return
	// }
	// fmt.Println("Course and category created successfully!")

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		fmt.Printf("Error listing courses: %v\n", err)
		return
	}
	for _, course := range courses {
		fmt.Printf("Course ID: %s, Name: %s, Description: %s, Price: %.2f, Category ID: %s, Category Name: %s\n",
			course.ID, course.Name, course.Description.String, course.Price, course.CategoryID, course.CategoryName)
	}
}
