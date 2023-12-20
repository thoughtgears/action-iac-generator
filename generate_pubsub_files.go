package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func (m *Module) generatePubSubTerraform(projectID string) error {
	path := "templates/modules/pubsub.tmpl"
	outputDir := "infrastructure"
	resourceName := sanitiseOutputFile(m.ResourceName)
	outputFilename := fmt.Sprintf("%s_pubsub.tf", resourceName)

	data := struct {
		ProjectID    string
		ResourceName string
		TopicName    string
	}{
		ProjectID:    projectID,
		ResourceName: resourceName,
		TopicName:    m.ResourceName,
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	outputFilePath := filepath.Join(outputDir, outputFilename)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(outputFile, data); err != nil {
		return err
	}

	log.Info().Msg("pub/sub terraform files generated successfully")

	return nil
}
