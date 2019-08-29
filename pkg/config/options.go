package config

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	TemplatePath flags.Filename `short:"t" long:"template" description:"Path of the template folder" default:"./"`
	OutputPath   flags.Filename `short:"o" long:"output" description:"Path of the output dir, if not specified the template will be generated in place" default:"./"`
}

func ParseCLIOption() (*Options, error) {
	var options Options
	_, err := flags.Parse(&options)
	if err != nil {
		return nil, err
	}
	return &options, nil
}
