package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// PubSub is the struct that contains the data for the pub/sub module.
type PubSub struct {
	TopicName string `yaml:"topic_name"`
}

func (p *PubSub) generateTerraform(projectID string) error {
	path := "templates/modules/pubsub.tmpl"
	outputDir := "infrastructure"
	outputFilename := fmt.Sprintf("%s_pubsub.tf", sanitiseOutputFile(p.TopicName))

	data := struct {
		ProjectID string
		TopicName string
	}{
		ProjectID: projectID,
		TopicName: p.TopicName,
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

	return nil
}
