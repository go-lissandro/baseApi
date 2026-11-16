package controllers

import (
	"goRest/models"
	"goRest/services"

	"github.com/gofiber/fiber/v2"
)

type CustommersService struct {
	Service services.CustommerService
}

type CustommersController interface {
	CreateCustommers(c *fiber.Ctx) error
	GetCustommers(c *fiber.Ctx) error
}

func (cs CustommersService) CreateCustommers(c *fiber.Ctx) error {
	var custommer models.Custommer
	if err := c.BodyParser(&custommer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid json",
			"message": err.Error(),
		})
	}

	err := cs.Service.CustommerInsert(custommer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Service unavailable.",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Criado com sucesso!",
	})
}
func (cs CustommersService) GetCustommers(c *fiber.Ctx) error {
	var custommers = cs.Service.CustommerGet()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get",
		"teste":   custommers,
	})

}

func NewCustommersController(service services.CustommerService) *CustommersService {
	return &CustommersService{service}
}
