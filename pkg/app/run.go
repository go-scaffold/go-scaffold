package app

import (
	"errors"
	"fmt"
	"log"
	"text/template"

	"github.com/go-scaffold/go-scaffold/pkg/config"
	"github.com/go-scaffold/go-scaffold/pkg/helpers"
	"github.com/go-scaffold/go-sdk/pkg/collectors"
	"github.com/go-scaffold/go-sdk/pkg/filters"
	"github.com/go-scaffold/go-sdk/pkg/pipeline"
	"github.com/go-scaffold/go-sdk/pkg/templateproviders"
	"github.com/go-scaffold/go-sdk/pkg/values"
)

var errHandler = log.Fatal

// Run starts the app
func Run(options *config.Options, funcMaps ...template.FuncMap) error {
	fileProvider := templateproviders.NewFileSystemProvider(string(options.TemplateDirPath()), filters.NewNoOpFilter())
	return RunWithFileProvider(options, fileProvider, funcMaps...)
}

// Run starts the app
func RunWithFileProvider(options *config.Options, fileProvider pipeline.TemplateProvider, funcMaps ...template.FuncMap) error {
	if options.TemplateRootPath == options.OutputPath {
		return errors.New("can't generate file in the input folder, please specify an output directory")
	}

	manifest, err := values.LoadYamlFilesWithPrefix("", options.ManifestPath())
	if err != nil {
		return errors.New(fmt.Sprintf("an error occurred while reading the manifest file: %s", err.Error()))
	}

	valuesPaths := make([]string, 0, len(options.Values)+1)
	valuesPaths = append(valuesPaths, options.ValuesPath())
	valuesPaths = append(valuesPaths, options.Values...)
	data, err := values.LoadYamlFilesWithPrefix("", valuesPaths...)
	if err != nil {
		return errors.New(fmt.Sprintf("error while loading data: %s", err.Error()))
	}

	pp, err := pipeline.NewBuilder().
		WithMetadata(manifest).
		WithData(data).
		WithFunctions(helpers.TemplateFunctions(funcMaps...)).
		WithTemplateProvider(fileProvider).
		WithCollector(collectors.NewFileWriterCollector(options.OutputPath, nil)).
		Build()
	if err != nil {
		return errors.New(fmt.Sprintf("error while building the processing pipeline: %s", err))
	}

	err = pp.Process()
	if err != nil {
		return errors.New(fmt.Sprintf("error while running the pipeline: %s", err.Error()))
	}
	return nil
}
