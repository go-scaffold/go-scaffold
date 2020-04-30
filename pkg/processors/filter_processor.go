package processors

import (
	"io"

	"github.com/pasdam/go-scaffold/pkg/filters"
)

type filterProcessor struct {
	filter        filters.Filter
	nextProcessor Processor
}

// NewFilterProcessor creates a new Processor that pass the file down to the
// specified nextProcessor when it matches the filter, it ignore the file
// otherwise
func NewFilterProcessor(filter filters.Filter, nextProcessor Processor) Processor {
	return &filterProcessor{
		filter:        filter,
		nextProcessor: nextProcessor,
	}
}

func (p *filterProcessor) ProcessFile(path string, reader io.Reader) error {
	if p.filter.Accept(path) {
		return p.nextProcessor.ProcessFile(path, reader)
	}
	return nil
}
