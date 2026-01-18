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
	templateProvider := templateproviders.NewFileSystemProvider(options.TemplateDirPath(), filters.NewNoOpFilter())

	namedTempaltesFilter, err := filters.NewPatternFilter(true, options.NamedTemplatesPattern)
	if err != nil {
		return fmt.Errorf("Invalid named templates pattern: %w", err)
	}
	namedTemplateProvider := templateproviders.NewFileSystemProvider(options.TemplateRootPath, namedTempaltesFilter)

	return RunWithCustomComponents(options, templateProvider, nil, namedTemplateProvider, funcMaps...)
}

// RunWithCustomComponents starts the app with custom components
func RunWithCustomComponents(options *config.Options, templateProvider pipeline.TemplateProvider, dataPreprocessor pipeline.DataPreprocessor, namedTemplateProvider pipeline.TemplateProvider, funcMaps ...template.FuncMap) error {
	if options.TemplateRootPath == options.OutputPath {
		return errors.New("can't generate file in the input folder, please specify an output directory")
	}

	loader := values.NewLoader()
	data, err := loader.LoadYAMLs(options.TemplateRootPath, options.Values)
	if err != nil {
		return fmt.Errorf("an error occurred while loading the data: %s", err.Error())
	}

	fileWriterOptions := collectors.FileWriterCollectorOptions{
		OutDir:           options.OutputPath,
		SkipUnchanged:    options.SkipUnchanged,
		CleanupUntracked: options.CleanupUntracked,
	}

	collector := collectors.NewSplitterCollector(
		collectors.NewFileWriterCollectorWithOpts(fileWriterOptions, nil),
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
		WithNamedTemplatesProvider(namedTemplateProvider).
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
