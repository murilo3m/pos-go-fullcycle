package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey`
	Name     string
	Products []Product
	gorm.Model
}

type Product struct {
	ID           int `gorm:"primaryKey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//Create Category
	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	//Create Product
	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      100,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	//Create SerialNumber
	// db.Create(&SerialNumber{
	// 	Number:    "1234",
	// 	ProductID: product.ID,
	// })

	//Belongs To
	// var products []Product
	// db.Preload("Category").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Has One
	// var products []Product
	// db.Preload("SerialNumber").Preload("Category").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Has Many
	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, ": ", product.SerialNumber.Number)
		}
	}
}
