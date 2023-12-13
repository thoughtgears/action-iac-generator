package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// generateBaseFiles generates the base terraform files. These files are the same for all services.
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

	return nil
}
