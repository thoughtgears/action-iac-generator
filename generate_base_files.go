package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// generateBaseFiles generates the base terraform files.
// It will generate base files that are the same for all services.
// It will read the templates from the templates directory and generate the files in the infrastructure
// directory. The infrastructure directory will be created if it does not exist.
// The base files are:
// - main.tf
// - provider.tf
func generateBaseFiles(config Config) error {
	templateDir := "templates"
	outputDir := "infrastructure"
	data := struct {
		TerraformVersion  string // Version of terraform to use
		Env               string // Environment for the service (e.g. "dev")
		StateBucket       string // Bucket name for the terraform state
		StateBucketPrefix string // Bucket prefix for the terraform state, generated from inputData
	}{
		TerraformVersion:  "1.5.3",
		Env:               config.Environment,
		StateBucket:       config.StateBucket,
		StateBucketPrefix: fmt.Sprintf("%s/%s/%s", config.Data.ProjectID, config.Data.Name, config.Environment),
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	files, err := os.ReadDir(templateDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		path := filepath.Join(templateDir, file.Name())

		if filepath.Ext(path) == ".tmpl" {
			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			outputFilename := strings.TrimSuffix(file.Name(), ".tmpl") + ".tf"
			outputFilePath := filepath.Join(outputDir, outputFilename)
			outputFile, err := os.Create(outputFilePath)
			if err != nil {
				return err
			}

			if err := tmpl.Execute(outputFile, data); err != nil {
				return err
			}
		}
	}

	log.Info().Msg("base terraform files generated successfully")

	return nil
}
