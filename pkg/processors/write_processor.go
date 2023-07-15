package processors

import (
	"io"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
	"github.com/pasdam/go-scaffold/pkg/core"
)

var ioutilxReaderToFile = ioutilx.ReaderToFile

type writeProcessor struct{}

// NewWriteProcessor creates a new instance of a Processor that write each the
// content from the reader to the specified path
func NewWriteProcessor() core.Processor {
	return &writeProcessor{}
}

func (p *writeProcessor) ProcessFile(filePath string, reader io.Reader) error {
	return ioutilxReaderToFile(reader, filePath)
}
