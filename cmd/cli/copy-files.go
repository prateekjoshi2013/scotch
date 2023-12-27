package main

import (
	"embed"
	"errors"
	"os"
)

//go:embed templates
var templateFS embed.FS

func copyFileFromTemplate(src, dst string) error {
	if fileExists(dst) {
		return errors.New("file already exists: " + dst)
	}
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

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
