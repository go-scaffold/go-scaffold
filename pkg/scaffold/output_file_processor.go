package scaffold

import (
	"io"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/processors"
)

type outputFileProcessor struct {
	outDir             string
	templateHelper     *TemplateHelper
	tempplateProcessor processors.Processor
	writeProcessor     processors.Processor
}

// NewOutputFileProcessor creates a new instance of a FileProcessor that process
// templates and creates the output files.
// THe variables to use for the template are in config.
func NewOutputFileProcessor(config interface{}, outDir string, templateHelper *TemplateHelper) processors.Processor {
	writeProcessor := processors.NewWriteProcessor()
	templateProcessor := processors.NewTemplateProcessor(config, writeProcessor)

	return &outputFileProcessor{
		outDir:             outDir,
		templateHelper:     templateHelper,
		tempplateProcessor: templateProcessor,
		writeProcessor:     writeProcessor,
	}
}

func (p *outputFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	outPath := filepath.Join(p.outDir, p.templateHelper.OutputFilePath(filePath))

	if p.templateHelper.Accept(filePath) {
		return p.tempplateProcessor.ProcessFile(outPath, reader)
	}
	return p.writeProcessor.ProcessFile(outPath, reader)
}
