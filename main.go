package main

import (
	r "goRest/goRest/app"
	"goRest/goRest/configs"
	"goRest/goRest/controllers"
	"goRest/goRest/repository"
	"goRest/goRest/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbClient := configs.ConnectPostgreSQL()

	custommersRepoDB := repository.NewCustommersDB(dbClient)
	custommersService := services.NewCustommersSevice(custommersRepoDB)
	custommerController := controllers.NewCustommersController(custommersService)

	r.SetupCustommersRoutes(app, custommerController)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
