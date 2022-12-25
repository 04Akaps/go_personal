package models

import (
	config "gorilla_mysql/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {

	var suc = db.NewRecord(b) // return true if primary key is blank
	if suc {
		db.Create(&b)
		return b
	} else {
		panic("Create Error : primary key already Existed")
	}
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(&book)
	return book
}
