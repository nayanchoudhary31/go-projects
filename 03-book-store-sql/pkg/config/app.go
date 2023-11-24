package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("msql", "iamnayan31:password/book_db?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("gorm Db connection error: ", err)
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
