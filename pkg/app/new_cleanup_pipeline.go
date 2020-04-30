package app

import (
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

func newCleanupPipeline(srcDir string) (processors.Processor, error) {
	filter, err := filters.NewPatternFilter(true, "\\.go-scaffold(/.*)?", "\\.*\\.tpl")
	if err != nil {
		return nil, err
	}
	removeProcessor := processors.NewRemoveFilesProcessor(srcDir)
	return processors.NewFilterProcessor(filter, removeProcessor), nil
}
