package app

import (
	"goRest/controllers"
	middlewre "goRest/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCustommersRoutes(c *fiber.App, cs controllers.CustommersController) {

	route := c.Group("/api")
	v1 := route.Group("/v1")

	v1.Post("/custommers", cs.CreateCustommers)

	v1.Group("/admin", middlewre.RoleValidationMiddleware(middlewre.RoleAdmin))

	v1.Get("/admin/custommers", cs.GetCustommers)
	v1.Get("/admin", func(c *fiber.Ctx) error {
		role := c.Locals("userRole").(middlewre.Role)
		return c.JSON(fiber.Map{
			"message": "Bem-vindo admin",
			"role":    role,
		})
	})

}
