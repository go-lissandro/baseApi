package controllers

import (
	"goRest/handler"
	middlewre "goRest/middleware"
	"goRest/models"
	"goRest/services"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	Service services.UserService
}

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
	LoginHandler(c *fiber.Ctx) error
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario criado",
		"body":    user,
	})
}

func (us UserService) LoginHandler(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body request", "msg": err})
	}
	findUser, err := us.Service.UserFind(body.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Erro ao buscar usaurio", "msg": err})
	}

	if handler.CheckPassword(findUser.Password, body.Password) {
		token, err := middlewre.GenerateToken("1", middlewre.RoleAdmin)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao errar token", "msg": err})
		}
		return c.JSON(fiber.Map{"token": token})

	}

	return c.JSON(fiber.Map{"invalid": "Credenciais Invalida"})
}

func NewUserController(service services.UserService) *UserService {
	return &UserService{service}
}
