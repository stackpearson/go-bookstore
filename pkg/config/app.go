package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// whole point of this file is to return our DB varriable and allow other files to interact with the db
var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "sawyer@Sawyers-MBP:SRP1992fish!/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func getDB() *gorm.DB {
	return db
}
