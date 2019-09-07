package config

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	OutputPath   flags.Filename `short:"o" long:"output" description:"Path of the output dir, if not specified the template will be generated in place" default:"./"`
	RemoveSource bool           `short:"r" long:"remove-source" description:"Flag to indicate whether remove the template and config files, or not. This has effect only if the input and output folder are the same"`
	TemplatePath flags.Filename `short:"t" long:"template" description:"Path of the template folder" default:"./"`
}

func ParseCLIOption() (*Options, error) {
	var options Options
	_, err := flags.Parse(&options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}
