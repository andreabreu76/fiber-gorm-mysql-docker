package moviments

import (
	"ackfinance/config"
	"github.com/gofiber/fiber/v2"
)

type movimentsResult struct {
	ID      string `json:"id"`
	Account string `json:"account"`
	Date    string `json:"date"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

func Read(c *fiber.Ctx) error {
	var movimentsResult []movimentsResult

	query := "SELECT" +
		"mov.id," +
		"CONCAT(a.bank_name, '-', a.number, '-', a.digit) AS account," +
		"DATE_FORMAT(mov.updated_at, '%d/%m/%Y - %H:%m') as date," +
		"mt.name AS type," +
		"mov.value" +
		"FROM moviments as mov" +
		"JOIN moviment_types mt on mov.id_type = mt.id" +
		"JOIN accounts a on mov.id_account = a.id "

	result := config.Database.Raw(query)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Erro ao buscar movimentos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(movimentsResult)
}
