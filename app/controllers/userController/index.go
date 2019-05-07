package userController

import (
	"ORM-example/app/models"
	"ORM-example/database"
	"net/http"
)

type UserAllController struct {
}

func (t UserAllController) ValidateInputs(r *http.Request) bool {
	return true
}

func (t UserAllController) Handle(w http.ResponseWriter, r *http.Request) interface{} {
	var users []models.User

	database.DB.Find(&users)

	return users
}
