package routes

import (
	"ackfinance/controllers/moviments"
	"github.com/gofiber/fiber/v2"
)

func MovimentsRoute(route fiber.Router) {
	route.Get("/", moviments.Read)
	route.Post("/", moviments.Create)
	route.Put("/:id", moviments.Update)
	route.Delete("/:id", moviments.Delete)
}
