package scaffold

import (
	"io"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

type FileProcessor struct {
	config         interface{}
	outDir         string
	templatesOnly  bool
	templateHelper *TemplateHelper
}

func NewFileProcessor(config interface{}, outDir string, templateHelper *TemplateHelper, templatesOnly bool) *FileProcessor {
	return &FileProcessor{
		config:         config,
		outDir:         outDir,
		templatesOnly:  templatesOnly,
		templateHelper: templateHelper,
	}
}

func (self *FileProcessor) ProcessFile(filePath string, reader io.Reader) error {
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

	return iohelpers.WriteFile(reader, filepath.Join(self.outDir, filePath))
}