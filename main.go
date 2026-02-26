package main

import (
	r "goRest/app"
	"goRest/configs"
	"goRest/controllers"
	"goRest/repository"
	"goRest/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173/",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	dbClient := configs.ConnectPostgreSQL()

	custommersRepoDB := repository.NewCustommersDB(dbClient)
	custommersService := services.NewCustommersSevice(custommersRepoDB)
	custommerController := controllers.NewCustommersController(custommersService)

	usersRepoDB := repository.NewUsersDB(dbClient)
	usersService := services.NewUserSevice(usersRepoDB)
	usersController := controllers.NewUserController(usersService)

	r.SetupCustommersRoutes(app, custommerController, usersController)

	err := app.Listen(":3000")

	if err != nil {
		return
	}
}
