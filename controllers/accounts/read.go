package accounts

import (
	"ackfinance/config"
	"ackfinance/models"

	"github.com/gofiber/fiber/v2"
)

func Read(c *fiber.Ctx) error {
	var accounts models.Account

	query := "" +
		"SELECT " +
		"a.bank_name, " +
		"CONCAT(a.number, '-', a.digit) AS account, " +
		"a.created_at, " +
		"astatus.name AS status " +
		"FROM accounts a " +
		"JOIN account_statuses astatus on a.id_status = astatus.id " +
		"ORDER BY a.bank_name DESC"

	result := config.Database.Raw(query).Scan(&accounts)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao buscar contas",
		})
	}

	return c.Status(fiber.StatusOK).JSON(accounts)
}
