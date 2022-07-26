package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "thneves:simplepass/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = database
}

func getDB() *gorm.DB {
	return db
}

// the whole poiunt is to return a db file
