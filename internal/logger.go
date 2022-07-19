package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func provideLogger(config *Config) *zerolog.Logger {
	logLevel, err := zerolog.ParseLevel(config.LogLevel)

	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = time.RFC3339

	if config.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: "15:04:05",
		})

		log.Logger = log.Logger.With().Caller().Logger()
	}

	log.Logger = log.Logger.With().Str("version", "alpha").Logger()

	return &log.Logger
}
