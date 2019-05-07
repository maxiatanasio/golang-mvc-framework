package userController

import (
	"ORM-example/app/models"
	"ORM-example/database"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserDeleteController struct {
}

func (UserDeleteController) ValidateInputs(r *http.Request) bool {
	return true
}

func (UserDeleteController) Handle(w http.ResponseWriter, r *http.Request) interface{} {
	vars := mux.Vars(r)
	name := vars["name"]

	var user models.User

	database.DB.Where("name = ?", name).First(&user)
	database.DB.Delete(&user)

	return map[string]string{
		"message": fmt.Sprintf("User #%v deleted", user.ID),
	}
}
