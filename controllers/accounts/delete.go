package accounts

import (
	"ackfinance/config"
	"ackfinance/models"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {

	var accounts models.Account

	var IDParam = c.Params("id")
	if IDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID não informado",
		})
	}

	result := config.Database.Where("id = ?", IDParam).Delete(&accounts)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao deletar conta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Conta deletada com sucesso",
	})

}
