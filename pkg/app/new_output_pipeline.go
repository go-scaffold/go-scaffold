package app

import (
	"text/template"

	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/helpers"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string, funcMap template.FuncMap) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	funcMap = helpers.TemplateFunctions(funcMap)
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
