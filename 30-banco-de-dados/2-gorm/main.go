package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product2 struct {
	ID    int `gorm:primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3303)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product2{})

	// db.Create(&Product2{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// var product Product2
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Notebook")
	// fmt.Println(product)

	// var products2 []Product2
	// db.Find(&products2)
	// for _, product := range products2 {
	// 	fmt.Println(product)
	// }

	// var products2 []Product2
	// db.Limit(2).Offset(2).Find(&products2)
	// for _, product := range products2 {
	// 	fmt.Println(product)
	// }

	// var products []Product2
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product2
	// db.Where("price LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	var p Product2
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product2
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	db.Delete(&p2, 1)
}
