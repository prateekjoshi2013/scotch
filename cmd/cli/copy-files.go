package main

import (
	"embed"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(src, dst string) error {
	// TODO check to ensure file does not already exist
	data, err := templateFS.ReadFile(src)
	if err != nil {
		return err
	}
	err = copyDataToFile(data, dst)
	if err != nil {
		return err
	}
	return nil
}

func copyDataToFile(data []byte, dst string) error {
	err := os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
