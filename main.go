package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	suiteSuffix = "_suite_test.go"
	fileSuffix  = "_test.go"
)

const (
	suiteTemplatePath = "./templates/templates_suite_test.go"
	fileTemplatePath  = "./templates/templates_test.go"
)

func main() {
	// Get the base name of the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Getwd():", err)
		return
	}
	folderName := filepath.Base(cwd)

	// Check if the suite file exists
	_, err = os.Stat(folderName + suiteSuffix)
	if err != nil {
		// Create SuiteFile if not exist
		err = cpFile(suiteTemplatePath, folderName+suiteSuffix, folderName)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(os.Args)

	fileName := folderName + fileSuffix
	if len(os.Args) >= 2 {
		fileName = os.Args[1] + fileSuffix
	}
	err = cpFile(fileTemplatePath, fileName, folderName)
}

const templates = "templates"

func cpFile(srcPath, destName, folderName string) error {
	// Read Template
	content, err := os.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}

	// Update Content
	modifiedContent := strings.Replace(string(content), templates, folderName, -1)

	// Write the modified content to the destination file
	err = os.WriteFile(destName, []byte(modifiedContent), 0644)
	if err != nil {
		return fmt.Errorf("os.Write: %w", err)
	}

	return nil
}
