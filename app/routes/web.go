package routes

import (
	"ORM-example/app/controllers"
	"ORM-example/app/controllers/userController"
	"ORM-example/gowitch"
)

var WebRoutes []gowitch.Route

func init() {

	WebRoutes = []gowitch.Route{
		gowitch.Route{
			Path:       "/user/{name}/{email}",
			Method:     "POST",
			Controller: userController.UserStoreController{},
		},
		gowitch.Route{
			Path:       "/users",
			Method:     "GET",
			Controller: userController.UserAllController{},
		},
		gowitch.Route{
			Path:       "/user/{name}",
			Method:     "DELETE",
			Controller: userController.UserDeleteController{},
		},
		gowitch.Route{
			Path:       "/user/{name}/{email}",
			Method:     "PUT",
			Controller: userController.UserUpdateController{},
		},
		gowitch.Route{
			Path:       "/test",
			Method:     "GET",
			Controller: controllers.TestController{},
		},
	}
}
