package internal

import (
	"github.com/goava/di"
	"github.com/rs/zerolog"
	"go-rest-api/internal/server"
)

func createContainer() (*di.Container, error) {
	return di.New(
		di.Invoke(LoadConfig),
		di.Provide(provideLogger),
		di.Provide(provideWebServer),
	)
}

func provideWebServer(config *Config, log *zerolog.Logger) *server.WebServer {
	return server.New(config, log)
}
