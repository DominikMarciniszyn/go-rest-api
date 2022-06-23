package main

import (
	"go-rest-api/internal/database/config"
	"go-rest-api/internal/routes"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	config.Connect()

	app.Get("/ping", routes.Ping)

	log.Fatal(app.Listen(":3000"))
}
