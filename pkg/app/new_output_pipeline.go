package app

import (
	"github.com/Masterminds/sprig/v3"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/helpers/collections"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newOutputPipeline(config interface{}, outDir string) (core.Processor, error) {
	filter := filters.NewNoOpFilter()
	funcMap := sprig.TxtFuncMap()
	funcMap["sequence"] = collections.Sequence
	outProcessor := processors.NewOutputFileProcessor(config, outDir, funcMap)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
