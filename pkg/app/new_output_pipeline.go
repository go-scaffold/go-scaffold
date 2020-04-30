package app

import (
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

func newOutputPipeline(inPlace bool, config interface{}, outDir string, templateHelper *scaffold.TemplateHelper) (processors.Processor, error) {
	var filter filters.Filter
	var err error
	if inPlace {
		filter, err = filters.NewPatternFilter(true, "\\.*\\.tpl")
	} else {
		filter, err = filters.NewPatternFilter(false, "\\.(go-scaffold|git)("+filepath.FromSlash("/")+".*)?")
	}

	if err != nil {
		return nil, err
	}

	outProcessor := scaffold.NewOutputFileProcessor(config, outDir, templateHelper)
	return processors.NewFilterProcessor(filter, outProcessor), nil
}
