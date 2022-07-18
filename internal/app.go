package internal

import (
	"github.com/rs/zerolog"
	"go-rest-api/internal/config"
	"go-rest-api/internal/server"
	"os"
	"os/signal"
)

type Container struct {
	Name string `db:"name"`
}

func Run() error {
	container, err := createContainer()

	if err != nil {
		return err
	}

	return container.Invoke(execute)
}

func execute(config *config.Config, webServer *server.WebServer, log *zerolog.Logger) (executeError error) {
	var err error

	log.Info().Msg("Starting container...")
	log.Debug().Interface("config", config).Msg("Loaded configuration...")

	stop := make(chan os.Signal, 1)

	if err != nil {
		return err
	}

	webServer.StartServerAsync(func(err error) {
		executeError = err
		stop <- os.Interrupt
	})

	signal.Notify(stop, os.Interrupt)
	<-stop

	return
}
