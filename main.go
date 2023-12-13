package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	log zerolog.Logger
)

func init() {
	log = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).Level(zerolog.InfoLevel).With().Timestamp().Caller().Logger()
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
	}

	fmt.Println(config)
}
