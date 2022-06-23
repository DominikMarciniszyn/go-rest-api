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
	app.Get("/orders", routes.GetOrders)
	app.Get("/orders/:id", routes.GetOrder)
	app.Post("/orders", routes.CreateOrder)
	app.Put("/orders/:id", routes.UpdateOrder)
	app.Delete("/orders/:id", routes.RemoveOrder)

	log.Fatal(app.Listen(":3000"))
}
