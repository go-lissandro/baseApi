package app

import (
	"goRest/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCustommersRoutes(c *fiber.App, cs controllers.CustommersController) {
	route := c.Group("/api")

	v1 := route.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Teste")
	})

	v1.Post("/custommers", cs.CreateCustommers)
	v1.Get("/custommers", cs.GetCustommers)
}
