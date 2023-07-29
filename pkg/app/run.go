package app

import (
	"log"

	"github.com/Masterminds/sprig"
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
func Run(options *config.Options) {
	fileProvider := templateproviders.NewFileSystemProvider(string(options.TemplateDirPath()), filters.NewNoOpFilter())
	RunWithFileProvider(options, fileProvider)
}

// Run starts the app
func RunWithFileProvider(options *config.Options, fileProvider pipeline.TemplateProvider) {
	if options.TemplateRootPath == options.OutputPath {
		log.Fatal("Can't generate file in the input folder, please specify an output directory")
		return
	}

	manifest, err := values.LoadYamlFilesWithPrefix("", options.ManifestPath())
	if err != nil {
		log.Fatal("An error occurred while reading the manifest file: ", err.Error())
		return
	}

	valuesPaths := make([]string, 0, len(options.Values)+1)
	valuesPaths = append(valuesPaths, options.ValuesPath())
	valuesPaths = append(valuesPaths, options.Values...)
	data, err := values.LoadYamlFilesWithPrefix("", valuesPaths...)
	if err != nil {
		log.Fatal("error while loading data. ", err)
		return
	}

	pp, err := pipeline.NewBuilder().
		WithMetadata(manifest).
		WithData(data).
		WithFunctions(helpers.TemplateFunctions(sprig.FuncMap())).
		WithTemplateProvider(fileProvider).
		WithCollector(collectors.NewFileWriterCollector(options.OutputPath, nil)).
		Build()
	if err != nil {
		log.Fatal("error while building the processing pipeline. ", err)
		return
	}

	err = pp.Process()
	if err != nil {
		log.Fatal("error while running the pipeline. ", err)
		return
	}
}
