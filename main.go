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

	if len(config.Data.Modules) > 0 {
		for _, module := range config.Data.Modules {
			switch module.Type {
			case "pubsub":
				if err := module.generatePubSubTerraform(config.Data.ProjectID); err != nil {
					log.Fatal().Err(err).Msg("failed to generate pubsub terraform files")
				}
			case "cloud-run":
				if err := module.generateCloudRunTerraform(config.Data.ProjectID, config.Data.Region); err != nil {
					log.Fatal().Err(err).Msg("failed to generate cloud run terraform files")
				}
			}
		}
	}
}
