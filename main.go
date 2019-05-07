package main

import (
	"ORM-example/app/middlewares"
	"ORM-example/app/models"
	"ORM-example/app/routes"
	"ORM-example/database"
	"ORM-example/gowitch"
	"fmt"
)

func handleRequests() {
	gowitch.RegisterMiddleware(middlewares.JSONMiddleware)
	for _, webRoute := range routes.WebRoutes {
		gowitch.RegisterRoute(webRoute)
	}
	gowitch.Init()
}

func main() {
	fmt.Println("Go ORM Tutorial")

	db := database.Connect()

	defer db.Close()

	database.RegisterEntity(models.User{})
	database.RegisterEntity(models.Pet{})

	database.LoadEntities()

	handleRequests()
}
