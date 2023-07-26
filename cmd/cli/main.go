package main

import (
	"fmt"
	"os"

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
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

	options := &config.Options{
		OutputPath:       args[1],
		TemplateRootPath: args[0],
		Values:           values,
	}

	app.Run(options)
}
