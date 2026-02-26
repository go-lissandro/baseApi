package controllers

import (
	"goRest/models"
	"goRest/services"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	Service services.UserService
}

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
}

func (us UserService) RegisterUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid json",
			"message": err.Error(),
		})
	}

	err := us.Service.UserInsert(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Service unavailable.",
			"message": err.Error(),
			"teste":   user,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario criado",
		"body":    user,
	})
}

func NewUserController(service services.UserService) *UserService {
	return &UserService{service}
}
