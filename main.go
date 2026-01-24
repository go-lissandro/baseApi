package main

import (
	r "goRest/app"
	"goRest/configs"
	"goRest/controllers"
	"goRest/repository"
	"goRest/services"

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
