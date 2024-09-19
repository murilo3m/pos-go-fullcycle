package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//Insert
	/*db.Create(&Product{
		Name:  "notebook",
		Price: 100.00,
	})

	products := []Product{
		{Name: "mouse", Price: 200.00},
		{Name: "teclado", Price: 200.00},
	}

	db.Create(&products)*/

	//Select - Específico
	/*var product Product
	db.First(&product, 1)
	fmt.Println(product)

	db.First(&product, "name = ?", "mouse")
	fmt.Println(product)*/

	//Select - Sem where
	/*var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}*/

	//Select - Com offset
	/*var products []Product
	db.Limit(2).Offset(0).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}*/

	//Select - Where
	/*var products []Product
	db.Where("price > ?", 100).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}*/

	//Select - Where (Like)
	/*var products []Product
	db.Where("name LIKE ?", "%book%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}*/

	//Update
	/*var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)*/

	//Delete
	/*var p Product
	db.First(&p, 1)
	db.Delete(&p)*/

	//Usando recursos do gorm.Model
	/*db.Create(&Product{
		Name:  "A",
		Price: 100,
	})*/

	/*var p Product
	db.First(&p, 5)
	p.Name = "New A"
	db.Save(&p)*/

	var p Product
	db.First(&p, 5)
	db.Delete(&p)

}
