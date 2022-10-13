package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func DbConnection() {
	dsn := "class:class@tcp(127.0.0.1:33061)/ht_students?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error while opening database connection :", err.Error())
		panic(err)
	}
	db = d
}

func GetDbConnection() *gorm.DB {
	return db
}
