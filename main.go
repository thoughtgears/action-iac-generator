package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger
)

func init() {
	log = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).Level(zerolog.InfoLevel).With().Timestamp().Logger()
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
	}

	if err := generateBaseFiles(config); err != nil {
		log.Fatal().Err(err).Msg("failed to generate base terraform files")
	}

	if config.Data.Modules != nil {
		if err := generateDynamicFiles(config); err != nil {
			log.Fatal().Err(err).Msg("failed to generate dynamic terraform files")
		}
	}
}
