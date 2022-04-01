package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sawyer/go-bookstore/pkg/config"
)

var db *gorm.DB

//structs are based off models and helps us store something in our DB
type Book struct {
	gorm.Model
	Name        string `gorm:""json: "name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

//helps us initialize the DB
func init() {
	config.Connect()
	db = config.GetDB()
	//this is just auto migrating the db with an empty book
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
