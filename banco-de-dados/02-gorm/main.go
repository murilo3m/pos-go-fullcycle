package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{
		Name:  "notebook",
		Price: 100.00,
	})

	products := []Product{
		{Name: "mouse", Price: 200.00},
		{Name: "teclado", Price: 200.00},
	}

	db.Create(&products)
}
