package app

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/iancoleman/strcase"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/helpers/collections"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string, funcMap template.FuncMap) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	sprigFuncMap := sprig.TxtFuncMap()
	sprigFuncMap["camelcase"] = strcase.ToCamel
	sprigFuncMap["replace"] = func(target string, replacement string, source string) string {
		return strings.ReplaceAll(source, target, replacement)
	}
	sprigFuncMap["sequence"] = collections.Sequence
	funcMap = mergeFuncMaps(sprigFuncMap, funcMap)
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
