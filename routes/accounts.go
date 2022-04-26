package routes

import (
	"ackfinance/controllers/accounts"
	"github.com/gofiber/fiber/v2"
)

func AccountsRoute(route fiber.Router) {
	route.Get("/", accounts.Read)
	route.Post("/", accounts.Create)
	route.Put("/:id", accounts.Update)
	route.Delete("/:id", accounts.Delete)
}
