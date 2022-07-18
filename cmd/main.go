package main

import (
	"github.com/rs/zerolog/log"
	"go-rest-api/internal"
	"os"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Err(err).Msg("Something went terribly wrong!")
		os.Exit(1)
	}

	os.Exit(0)
}
