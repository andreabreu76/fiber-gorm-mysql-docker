package moviments

import (
	"ackfinance/config"
	"ackfinance/models"
	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	var moviments models.Moviments

	IDParam := c.Params("id")
	if IDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID is required",
		})
	}

	result := config.Database.Where("id = ?", IDParam).Delete(&moviments)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Moviment not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Moviment deleted",
	})

}
