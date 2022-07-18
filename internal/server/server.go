package server

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"go-rest-api/internal"
	"go-rest-api/internal/routes"
	"log"
	"os"
)

type WebServer struct {
	config *internal.Config
	server *fiber.App
	log    *zerolog.Logger
	stop   chan os.Signal
}

func New(config *internal.Config, log *zerolog.Logger) *WebServer {
	app := fiber.New()

	return &WebServer{
		config: config,
		server: app,
		log:    log,
		stop:   make(chan os.Signal, 1),
	}
}

func (s *WebServer) DefineRoutes() {
	api := s.server.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/ping", routes.Ping)
	v1.Get("/orders", routes.GetOrders)
	v1.Get("/orders/:id", routes.GetOrder)
	v1.Post("/orders", routes.CreateOrder)
	v1.Put("/orders/:id", routes.UpdateOrder)
	v1.Delete("/orders/:id", routes.RemoveOrder)

	log.Fatal(s.server.Listen(viper.GetInt("port")))
}

func (s *WebServer) StartServerAsync(errCallback func(err error)) {
	go func() {
		s.log.Info().Msgf("Starting server at port: %d", s.config.Port)

		err := s.server.Listen(s.config.Port)

		s.log.Info().Msg("Stopping server...")
		s.stop <- os.Interrupt

		if err != nil && errCallback != nil {
			errCallback(err)
		}
	}()
}

func Shutdown(s *WebServer) error {
	err := s.server.Shutdown()
	<-s.stop

	return err
}
