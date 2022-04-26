package main

import (
	"ackfinance/config"
	"ackfinance/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found... copying .env.example")
	}

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true})

	config.Default(app)
	config.Connect()
	routes.Setup(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
