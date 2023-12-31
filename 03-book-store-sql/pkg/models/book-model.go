package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nayanchoudhary31/03-book-store-sql/pkg/config"
)

var db *gorm.DB

// Book Struct
type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

// Database Connection
func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}
