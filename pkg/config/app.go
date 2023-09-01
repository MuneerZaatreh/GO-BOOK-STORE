package config
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	db * gorm.DB
)
func Connect() {
    dsn := "root:@tcp(localhost:3306)/BookStore?charset=utf8&parseTime=True&loc=Local"
    d, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }
    db = d
}

func GetDb() *gorm.DB {
	return db;
}