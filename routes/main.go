package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"projectName": "Basics Finance!",
			"version":     "1.0.0",
			"author":      "Andre Abreu",
			"description": "A simple API to manage your finances with Air for DEV env Help, Gorm, Fiber and MySQL.",
		})
	})

	api := app.Group("/api")
	AccountsRoute(api.Group("/accounts"))
	MovimentsRoute(api.Group("/moviments"))
}
