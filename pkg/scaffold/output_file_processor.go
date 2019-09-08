package scaffold

import (
	"io"
	"log"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

type OutputFileProcessor struct {
	config         interface{}
	outDir         string
	templatesOnly  bool
	templateHelper *TemplateHelper
}

func NewOutputFileProcessor(config interface{}, outDir string, templateHelper *TemplateHelper, templatesOnly bool) FileProcessor {
	return &OutputFileProcessor{
		config:         config,
		outDir:         outDir,
		templatesOnly:  templatesOnly,
		templateHelper: templateHelper,
	}
}

func (self *OutputFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	var err error
	if self.templateHelper.Accept(filePath) {
		reader, err = ProcessTemplate(reader, self.config)
		if err != nil {
			return err
		}
		filePath = self.templateHelper.OutputFilePath(filePath)

	} else if self.templatesOnly {
		// ingore normal files
		return nil
	}

	log.Printf("Writing file %s\n", filepath.Join(self.outDir, filePath))

	return iohelpers.WriteFile(reader, filepath.Join(self.outDir, filePath))
}
