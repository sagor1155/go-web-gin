package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name   string
	Author string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!!")
	}
	fmt.Println("Database connected successfully!!")
	db.AutoMigrate(&Book{})

	// create
	db.Create(&Book{
		Name:   "Ocean 11",
		Author: "Mr Duke",
	})

	// read
	var firstBook Book
	db.First(&firstBook, 1)
	fmt.Println(firstBook.Name)
	fmt.Println(firstBook.Author)
	fmt.Println(firstBook.CreatedAt)
	fmt.Println(firstBook.UpdatedAt)

	var myBook Book
	db.First(&myBook, "name = ?", "Ocean 11")
	fmt.Print(myBook.Name)

	// update
	// db.Model(&myBook).Update("Author", "marly duke").Where("ID", 1)
	db.Model(&myBook).Updates(&Book{Name: "Ocean 12", Author: "Joe Doe"}).Where("ID", 2)
}
