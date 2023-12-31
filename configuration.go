package main

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

// InputData is the struct that contains the data that will be read from the infrastructure
// yaml file and be used to generate the terraform files from the templates.
// The input data will also contain unexported field that will be used to store interpolated
// data based on the input data.
type InputData struct {
	Name      string    `yaml:"name"`              // Name of the service (e.g. "my-service")
	ProjectID string    `yaml:"project_id"`        // ProjectID for the service (e.g. "my-project")
	Region    string    `yaml:"region"`            // Region for the service (e.g. "us-central1")
	Modules   []*Module `yaml:"modules,omitempty"` // Modules for the service (e.g. "pubsub")
}

// Module is the struct that contains the data for the modules that will be used to generate
// the terraform files from the templates. The modules will be used to generate dynamic files
// that are specific to the service.
type Module struct {
	Type           string          `yaml:"type"`          // Type of module (e.g. "pubsub")
	ResourceName   string          `yaml:"resource_name"` // Name of resource, both from a resource and naming convention perspective (e.g. "topic-x")
	CloudRunModule *CloudRunModule `yaml:"cloud_run,omitempty"`
}

// Config is the struct that contains the data that will be read from the environment variables
// and the infrastructure yaml file.
type Config struct {
	FileName    string `envconfig:"FILE_NAME" default:"infrastructure.yaml"`
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
	StateBucket string `envconfig:"STATE_BUCKET" required:"true"`
	Data        InputData
}

// getConfig reads the environment variables and sets the values in the Config struct.
// It will also read the infrastructure yaml file and set the values in the Config struct.
// The Config struct will be used to generate the terraform files from the templates.
func getConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("input", &config); err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(config.FileName)
	if err != nil {
		return Config{}, err
	}

	if err := yaml.Unmarshal(file, &config.Data); err != nil {
		return Config{}, err
	}

	return config, nil
}
