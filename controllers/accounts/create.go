package accounts

import (
	"ackfinance/config"
	"ackfinance/models"

	"github.com/go-playground/validator/v10"
	//https://github.com/go-playground/validator/blob/master/_examples/simple/main.go

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	account := new(models.Account)

	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao criar conta",
		})
	}

	validate := validator.New()
	err := validate.Struct(account)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Erro ao criar conta",
			"errors":  err,
		})
	}

	result := config.Database.Create(&account)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao criar conta",
			"errors":  result.Error,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Conta criada com sucesso",
		"data":    account,
	})
}
