package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type CloudRunModule struct {
	Image  string `yaml:"image"` // Image for the service (e.g. "gcr.io/my-project/my-service:latest")
	Limits struct {
		CPU    string `yaml:"cpu"`    // CPU limit for the service (e.g. "1000m")
		Memory string `yaml:"memory"` // Memory limit for the service (e.g. "256Mi")
	} `yaml:"limits"`
}

func (m *Module) generateCloudRunTerraform(projectID, region string) error {
	path := "templates/modules/cloud_run.tmpl"
	outputDir := "infrastructure"
	resourceName := sanitiseOutputFile(m.ResourceName)
	outputFilename := fmt.Sprintf("%s_cloud_run.tf", resourceName)

	data := struct {
		ProjectID    string
		Location     string
		ResourceName string
		ServiceName  string
		Image        string
		CPU          string
		Memory       string
	}{
		ProjectID:    projectID,
		Location:     region,
		ResourceName: resourceName,
		ServiceName:  m.ResourceName,
		Image:        m.CloudRunModule.Image,
		CPU:          m.CloudRunModule.Limits.CPU,
		Memory:       m.CloudRunModule.Limits.Memory,
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
