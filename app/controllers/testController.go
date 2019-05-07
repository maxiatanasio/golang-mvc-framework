package controllers

import (
	"ORM-example/app/models"
	"ORM-example/database"
	"net/http"
)

type TestController struct {
	Name string
}

func (t TestController) ValidateInputs(r *http.Request) bool {
	return true
}

func (t TestController) Handle(w http.ResponseWriter, r *http.Request) interface{} {
	var users []models.User

	database.DB.Find(&users)

	return users
}
