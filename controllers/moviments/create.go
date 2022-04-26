package moviments

import (
	"ackfinance/config"
	"ackfinance/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	moviments := new(models.Moviments)

	if err := c.BodyParser(moviments); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(moviments)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao criar movimentação",
			"error":   err.Error(),
		})
	}

	result := config.Database.Create(&moviments)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao criar movimentação",
			"error":   result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Movimentação criada com sucesso",
		"data":    moviments,
	})
}
