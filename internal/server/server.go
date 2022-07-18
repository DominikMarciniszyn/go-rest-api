package server

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
	"go-rest-api/internal/routes"
	"os"
)

type WebServer struct {
	config *Config
	server *fiber.App
	log    *zerolog.Logger
	stop   chan os.Signal
}

func New(config *Config, log *zerolog.Logger) *WebServer {
	app := fiber.New()

	return &WebServer{
		config: config,
		server: app,
		log:    log,
		stop:   make(chan os.Signal, 1),
	}
}

func (s *WebServer) StartServerAsync(errCallback func(err error)) {
	go func() {
		s.log.Info().Msgf("Starting server at port: %d", s.config.Port)

		router := s.server.Group("/v1")
		router.Get("/ping", routes.Ping)
		router.Get("/orders", routes.GetOrders)
		router.Get("/orders/:id", routes.GetOrder)
		router.Post("/orders", routes.CreateOrder)
		router.Put("/orders/:id", routes.UpdateOrder)
		router.Delete("/orders/:id", routes.RemoveOrder)

		err := s.server.Listen(3344)

		s.log.Info().Msg("Stopping server...")
		s.stop <- os.Interrupt

		if err != nil && errCallback != nil {
			errCallback(err)
		}
	}()
}

func (s *WebServer) Shutdown() error {
	err := s.server.Shutdown()
	<-s.stop

	return err
}
