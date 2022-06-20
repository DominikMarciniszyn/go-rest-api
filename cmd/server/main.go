package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/ping", Ping)

	log.Fatal(app.Listen(":3000"))
}
