package models

import "github.com/jinzhu/gorm"

type Pet struct {
	gorm.Model
	Name   string
	Animal string
}
