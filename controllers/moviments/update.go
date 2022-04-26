package moviments

import (
	"ackfinance/config"
	"ackfinance/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	moviments := new(models.Moviments)

	IDParam := c.Params("id")
	if IDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID is required",
		})
	}

	result := config.Database.First(moviments, IDParam)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Erro ao buscar conta",
		})
	}

	if err := c.BodyParser(moviments); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	validate := validator.New()
	err := validate.Struct(moviments)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Erro ao validar movimento",
			"message": err.Error(),
		})
	}

	result = config.Database.Save(moviments)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao atualizar movimento",
		})
	}

	return c.Status(fiber.StatusOK).JSON(moviments)
}
