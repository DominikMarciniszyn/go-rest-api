package server

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
	"go-rest-api/internal/database/config"
	"go-rest-api/internal/routes"
	"log"
)

type WebServer struct {
	server *fiber.App
	log    *zerolog.Logger
	port   *int
}

func New(log *zerolog.Logger, port *int) *fiber.App {
	app := fiber.New()

	err := app.Listen(port)
	if err != nil {
		return nil
	}

	return nil
}

func StartServer() error {
	app := fiber.New()

	api := app.Group("/api")
	err := config.Connect()

	if err != nil {
		log.Fatal("Cannot connect with the database!")
	}

	v1 := api.Group("/v1")
	v1.Get("/ping", routes.Ping)
	v1.Get("/orders", routes.GetOrders)
	v1.Get("/orders/:id", routes.GetOrder)
	v1.Post("/orders", routes.CreateOrder)
	v1.Put("/orders/:id", routes.UpdateOrder)
	v1.Delete("/orders/:id", routes.RemoveOrder)

	log.Fatal(app.Listen(":3000"))

	return nil
}
