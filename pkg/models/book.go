package models
import (
	"github.com/jinzhu/gorm"
	"github.com/MuneerZaatreh/go-bookstore/pkg/config"
)
var db *gorm.DB 

type Book struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
}

func init (){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b * Book) CreateBook() * Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetBookById(Id int64) (*Book, *gorm.DB) {
    var getBook Book
    db := db.Where("ID=?", Id).Find(&getBook)
    return &getBook, db
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}



func DeleteBook (Id int64) *Book {
	var book Book
	db.Where("ID=?",Id).Delete(&book)
	return &book
}