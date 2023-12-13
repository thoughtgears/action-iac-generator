package main

import (
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

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			outputFilename := strings.TrimSuffix(filepath.Base(path), ".tmpl") + ".tf"
			outputFilePath := filepath.Join(outputDir, outputFilename)
			outputFile, err := os.Create(outputFilePath)
			if err != nil {
				return err
			}

			if err := tmpl.Execute(outputFile, config.Data); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	log.Info().Msg("base terraform files generated successfully")

	return nil
}
