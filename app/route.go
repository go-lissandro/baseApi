package app

import (
	"goRest/controllers"
	middlewre "goRest/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCustommersRoutes(c *fiber.App, cs controllers.CustommersController, us controllers.UserController) {

	route := c.Group("/api")
	v1 := route.Group("/v1")

	v1.Post("/create-user", us.RegisterUser)
	v1.Post("/login", us.LoginHandler)

	v1.Post("/custommers", cs.CreateCustommers)

	// O grupo precisa ser guardado: antes o retorno era descartado e as rotas
	// /admin ficavam registradas em v1, sem nenhuma autenticação.
	admin := v1.Group("/admin", middlewre.RoleValidationMiddleware(middlewre.RoleAdmin))

	admin.Get("/custommers", cs.GetCustommers)
	admin.Get("/", func(c *fiber.Ctx) error {
		role, _ := c.Locals(middlewre.RoleCtxKey).(middlewre.Role)
		return c.JSON(fiber.Map{
			"message": "Bem-vindo admin",
			"role":    role,
		})
	})

}
