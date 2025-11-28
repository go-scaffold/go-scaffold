package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-scaffold/go-scaffold/pkg/app"
	"github.com/go-scaffold/go-scaffold/pkg/config"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-scaffold",
		Short: "go-scaffold generates files and projects from templates",
	}

	generateCmd := &cobra.Command{
		Use:   "generate [flags] template-dir output-dir",
		Short: "generates files and projects from templates",
		Run:   generate,
		Args:  cobra.ExactArgs(2),
	}
	generateCmd.Flags().StringArrayP("values", "f", nil, "overriding values file")
	generateCmd.Flags().Bool("skip-unchanged", false, "skip writing files that would be unchanged")
	generateCmd.Flags().Bool("cleanup-untracked", false, "remove untracked files that are no longer part of the template")
	rootCmd.AddCommand(generateCmd)

	createCmd := &cobra.Command{
		Use:   "create [name]",
		Short: "create initializes a new template project",
		Long:  `create initializes a new template project with the required file structure, including Manifest.yaml, values.yaml, and a templates directory.`,
		Run:   create,
		Args:  cobra.MaximumNArgs(1),
	}
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func create(cmd *cobra.Command, args []string) {
	var targetDir string
	if len(args) > 0 {
		targetDir = args[0]
	} else {
		// Check if current directory is empty
		isEmpty, err := isDirEmpty(".")
		if err != nil {
			fmt.Printf("Error checking directory: %v\n", err)
			os.Exit(1)
		}
		if !isEmpty {
			fmt.Println("Error: current directory is not empty. Please provide a name for the new template or use an empty directory.")
			os.Exit(1)
		}
		targetDir = "."
	}

	// If targetDir is not ".", create the subdirectory
	if targetDir != "." {
		if err := os.MkdirAll(targetDir, 0644); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", targetDir, err)
			os.Exit(1)
		}
	}

	// Create the required files
	manifestPath := filepath.Join(targetDir, "Manifest.yaml")
	valuesPath := filepath.Join(targetDir, "values.yaml")
	templatesDir := filepath.Join(targetDir, "templates")

	// Create Manifest.yaml
	manifestContent := `# Manifest.yaml defines the template configuration
version: "1.0"
name: my-template
description: A sample template
`
	if err := os.WriteFile(manifestPath, []byte(manifestContent), 0644); err != nil {
		fmt.Printf("Error creating Manifest.yaml: %v\n", err)
		os.Exit(1)
	}

	// Create values.yaml
	valuesContent := `# values.yaml contains default values for the template
name: my-app
image:
  repository: nginx
  tag: latest
replicaCount: 1
`
	if err := os.WriteFile(valuesPath, []byte(valuesContent), 0644); err != nil {
		fmt.Printf("Error creating values.yaml: %v\n", err)
		os.Exit(1)
	}

	// Create templates directory
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		fmt.Printf("Error creating templates directory: %v\n", err)
		os.Exit(1)
	}

	// Create a simple example template in the templates directory
	templatePath := filepath.Join(templatesDir, "example.yaml")
	templateContent := `# Example template file
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .name }}
data:
  app_name: {{ .name }}
  replica_count: "{{ .replicaCount }}"
`
	if err := os.WriteFile(templatePath, []byte(templateContent), 0644); err != nil {
		fmt.Printf("Error creating example template: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Template created successfully in %s/\n", targetDir)
}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Read only one entry
	if err == nil {
		return false, nil // Directory is not empty
	}
	if err == os.ErrNotExist || err == io.EOF {
		return true, nil // Directory is empty or doesn't exist
	}
	return false, err // Some other error occurred
}

func generate(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.PrintErr("Invalid number of command arguments")
		cmd.Help()
		return
	}

	values, err := cmd.Flags().GetStringArray("values")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	skipUnchanged, err := cmd.Flags().GetBool("skip-unchanged")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cleanupUntracked, err := cmd.Flags().GetBool("cleanup-untracked")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	options := &config.Options{
		OutputPath:       args[1],
		TemplateRootPath: args[0],
		Values:           values,
		SkipUnchanged:    skipUnchanged,
		CleanupUntracked: cleanupUntracked,
	}

	err = app.Run(options)
	if err != nil {
		log.Fatal(err)
	}
}
