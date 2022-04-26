package accounts

import (
	"ackfinance/config"
	"ackfinance/models"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	account := new(models.Account)

	IDParam := c.Params("id")
	if IDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID não informado",
		})
	}

	result := config.Database.First(&account, IDParam)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao buscar conta",
		})
	}

	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao atualizar conta",
		})
	}

	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Erro ao validar conta",
			"errors":  err,
		})
	}

	result = config.Database.Save(&account)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao atualizar conta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(account)
}
