package models

import (
	"github.com/kHamidulloh/Go-BookStore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, error) {
	var book Book
	result := db.Where("id = ?", ID).First(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func UpdateBook(book *Book) (*Book, error) {
	result := db.Save(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func DeleteBook(id int64) error {
	var book Book
	result := db.Where("id=?", id).Delete(&book)
	return result.Error
}
