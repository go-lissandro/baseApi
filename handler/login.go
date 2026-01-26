package handler

import (
	middlewre "goRest/middleware"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body request", "msg": err})
	}

	//mock
	if !CheckPassword("123456", body.Password) {
		token, _ := middlewre.GenerateToken("1", middlewre.RoleAdmin)
		return c.JSON(fiber.Map{"token": token})
	}

	//mock
	token, err := middlewre.GenerateToken("2", middlewre.RoleUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao errar token", "msg": err})
	}

	//mock

	return c.JSON(fiber.Map{"token": token, "role": "role"})
}
