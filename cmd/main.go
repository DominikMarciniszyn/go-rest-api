package main

import (
	"go-rest-api/internal/routes"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/ping", routes.Ping)

	log.Fatal(app.Listen(":3000"))
}
