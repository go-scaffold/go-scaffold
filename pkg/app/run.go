package app

import (
	"errors"
	"fmt"
	"text/template"

	"github.com/go-scaffold/go-scaffold/pkg/config"
	"github.com/go-scaffold/go-scaffold/pkg/helpers"
	"github.com/go-scaffold/go-sdk/v2/pkg/collectors"
	"github.com/go-scaffold/go-sdk/v2/pkg/filters"
	"github.com/go-scaffold/go-sdk/v2/pkg/pipeline"
	"github.com/go-scaffold/go-sdk/v2/pkg/templateproviders"
	"github.com/go-scaffold/go-sdk/v2/pkg/values"
)

// Run starts the app
func Run(options *config.Options, funcMaps ...template.FuncMap) error {
	fileProvider := templateproviders.NewFileSystemProvider(string(options.TemplateDirPath()), filters.NewNoOpFilter())
	return RunWithCustomComponents(options, fileProvider, nil, funcMaps...)
}

// Run starts the app
func RunWithCustomComponents(options *config.Options, templateProvider pipeline.TemplateProvider, dataPreprocessor pipeline.DataPreprocessor, funcMaps ...template.FuncMap) error {
	if options.TemplateRootPath == options.OutputPath {
		return errors.New("can't generate file in the input folder, please specify an output directory")
	}

	loader := values.NewLoader()
	data, err := loader.LoadYAMLs(options.TemplateRootPath, options.Values)
	if err != nil {
		return fmt.Errorf("an error occurred while loading the data: %s", err.Error())
	}

	collector := collectors.NewSplitterCollector(
		collectors.NewFileWriterCollector(options.OutputPath, nil),
	)

	customFuncMap := make(template.FuncMap)
	customFuncMap["fileHeader"] = collector.CreateHeaderWithName
	funcMaps = append(funcMaps, customFuncMap)

	pp, err := pipeline.
		NewPipelineBuilder().
		WithDataPreprocessor(dataPreprocessor).
		WithFunctions(helpers.TemplateFunctions(funcMaps...)).
		WithTemplateAwareFunctions(helpers.TemplateAwareFunctions()).
		WithTemplateProvider(templateProvider).
		WithCollector(collector).
		Build()
	if err != nil {
		return fmt.Errorf("error while building the processing pipeline: %s", err)
	}

	err = pp.Process(data)
	if err != nil {
		return fmt.Errorf("error while running the pipeline: %s", err.Error())
	}
	return nil
}
