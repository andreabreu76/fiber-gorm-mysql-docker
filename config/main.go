package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Default(app *fiber.App) {

	app.Use(cors.New())
	app.Use(logger.New())

}
