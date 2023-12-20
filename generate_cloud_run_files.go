package main

type CloudRun struct {
	ServiceName string `yaml:"service_name"`
}

func (c *CloudRun) generateTerraform() error {
	return nil
}
