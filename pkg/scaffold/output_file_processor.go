package scaffold

import (
	"io"
	"log"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/templates"
)

type outputFileProcessor struct {
	config         interface{}
	outDir         string
	templatesOnly  bool
	templateHelper *TemplateHelper
}

// NewOutputFileProcessor creates a new instance of a FileProcessor that process templates and creates the output files.
// THe variables to use for the template are in config. It can be used to process template files only.
func NewOutputFileProcessor(config interface{}, outDir string, templateHelper *TemplateHelper, templatesOnly bool) FileProcessor {
	return &outputFileProcessor{
		config:         config,
		outDir:         outDir,
		templatesOnly:  templatesOnly,
		templateHelper: templateHelper,
	}
}

func (p *outputFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	var err error
	if p.templateHelper.Accept(filePath) {
		reader, err = templates.ProcessTemplate(reader, p.config)
		if err != nil {
			return err
		}
		filePath = p.templateHelper.OutputFilePath(filePath)

	} else if p.templatesOnly {
		// ingore normal files
		return nil
	}

	log.Printf("Writing file %s\n", filepath.Join(p.outDir, filePath))

	return iohelpers.WriteFile(reader, filepath.Join(p.outDir, filePath))
}
