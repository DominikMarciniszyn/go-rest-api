package main

import (
	"go-rest-api/internal/routes"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/ping", routes.Ping)

	log.Fatal(app.Listen(":3000"))
}
