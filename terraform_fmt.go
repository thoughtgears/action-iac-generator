package main

import "os/exec"

func terraformFmt() error {
	terraformDir := "./infrastructure"

	// Create the command to run 'terraform fmt'
	cmd := exec.Command("terraform", "fmt", terraformDir)

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	log.Info().Msg("terraform fmt run successfully")

	return nil
}
