package processors

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

type removeFilesProcessor struct {
	rootDir string
}

// NewRemoveFilesProcessor returns a processor that removes files
func NewRemoveFilesProcessor(rootDir string) Processor {
	return &removeFilesProcessor{
		rootDir: rootDir,
	}
}

func (p *removeFilesProcessor) ProcessFile(path string, _ io.Reader) error {
	if len(p.rootDir) > 0 {
		path = filepath.Join(p.rootDir, path)
	}

	err := os.Remove(path)
	if err != nil {
		return err
	}

	parent := filepath.Dir(path)
	return iohelpers.RemoveDirIfEmpty(parent)
}
