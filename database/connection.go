package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error
var entities []interface{}

func Connect() *gorm.DB {
	DB, err = gorm.Open("mysql", "root@/simpleUsers?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	return DB
}

func RegisterEntity(entity interface{}) {
	entities = append(entities, entity)
}

func LoadEntities() {
	for _, entity := range entities {
		DB.AutoMigrate(entity)
	}
}
