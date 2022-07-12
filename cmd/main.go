package main

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog/log"
	"go-rest-api/internal"
	"go-rest-api/internal/config"
	"os"
	"os/signal"
)

func main() {
	config.InitServiceConfig()

	if err := internal.Invoke(startHttpServer); err != nil {
		panic(err)
	}
}

func startHttpServer(server *fiber.App) error {
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Trace().Msg("Received SIGINT")

		err := server.Shutdown()

		if err != nil {
			return
		}

		log.Trace().Msg("Shutdown gracefully")

		done <- true
	}()

	<-done
	return nil
}
