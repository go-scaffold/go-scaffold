package processors

import (
	"io"
)

type pipeline struct {
	procs []Processor
}

// NewPipeline creates a pipeline of Processor, called in the same order as they
// are passed in the constructor
func NewPipeline(procs ...Processor) Processor {
	return &pipeline{
		procs: procs,
	}
}

func (p *pipeline) ProcessFile(path string, reader io.Reader) error {
	for _, proc := range p.procs {
		err := proc.ProcessFile(path, reader)
		if err != nil {
			return err
		}
	}
	return nil
}
