package main

import (
	"os"
	"path/filepath"
	"text/template"
)

// generateDynamicFiles generates the base terraform files.
// It will generate dynamic files that are specific to the service.
// It will read the templates from the templates/modules directory and generate the files in the infrastructure
// directory. The infrastructure directory will be created if it does not exist.
// The dynamic files are:
// - pubsub.tf
func generateDynamicFiles(config Config) error {
	templateDir := "templates/modules"
	outputDir := "infrastructure"

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	for _, module := range config.Data.Modules {
		templatePath := filepath.Join(templateDir, module.Name+".tmpl")

		if _, err := os.Stat(templatePath); os.IsNotExist(err) {
			continue
		}

		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			return err
		}

		outputFilename := module.Name + ".tf"
		outputFilePath := filepath.Join(outputDir, outputFilename)
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			return err
		}

		if err := tmpl.Execute(outputFile, config.Data); err != nil {
			return err
		}
	}

	log.Info().Msg("dynamic terraform files generated successfully")

	return nil
}
