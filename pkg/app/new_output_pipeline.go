package app

import (
	"strings"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/helpers/collections"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	funcMap := sprig.TxtFuncMap()
	funcMap["camelcase"] = strcase.ToCamel
	funcMap["replace"] = func(target string, replacement string, source string) string {
		return strings.ReplaceAll(source, target, replacement)
	}
	funcMap["sequence"] = collections.Sequence
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
