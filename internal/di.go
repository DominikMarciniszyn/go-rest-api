package internal

import (
	"errors"
	"github.com/goava/di"
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-rest-api/internal/config"
	"os"
)

var (
	Container *di.Container
)

func Invoke(invocation di.Invocation) error {
	return Container.Invoke(invocation)
}

func init() {
	var err error

	Container, err = di.New(di.Provide(provideLogger))

	if err != nil {
		panic(err)
	}
}

func provideLogger(config *config.Config) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(config.LogLevel)

	if err != nil {
		return nil, errors.New("invalid value for log level! Check the configuration")
	}

	logger := log.Logger.Level(level)

	if config.Pretty {
		logger = logger.Output(zerolog.ConsoleWriter{
			Out: os.Stderr,
		})
	}

	return &logger, nil
}

func provideWebServer(log *zerolog.Logger, config *config.Config) *fiber.App {
	return fiber.New()
}
