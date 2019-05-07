package userController

import (
	"ORM-example/app/models"
	"ORM-example/database"
	"github.com/gorilla/mux"
	"net/http"
)

type UserStoreController struct {
}

func (u UserStoreController) ValidateInputs(r *http.Request) bool {
	return true
}

func (u UserStoreController) Handle(w http.ResponseWriter, r *http.Request) interface{} {
	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	user := models.User{
		Name:  name,
		Email: email,
	}

	database.DB.Create(&user)

	return user
}
