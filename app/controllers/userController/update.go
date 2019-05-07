package userController

import (
	"ORM-example/app/models"
	"ORM-example/database"
	"github.com/gorilla/mux"
	"net/http"
)

type UserUpdateController struct {
}

func (UserUpdateController) ValidateInputs(r *http.Request) bool {
	return true
}

func (UserUpdateController) Handle(w http.ResponseWriter, r *http.Request) interface{} {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user models.User

	database.DB.Where("name = ?", name).First(&user)

	user.Email = email
	user.Model.DeletedAt = nil

	database.DB.Save(&user)

	return user
}
