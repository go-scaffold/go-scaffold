package providers

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pasdam/files-index/pkg/filesindex"
	"github.com/pasdam/go-scaffold/pkg/core"
)

var open = os.Open

type fileSystemProvider struct {
	indexer *filesindex.Indexer
}

// NewFileSystemProvider creates a new instance of a FileProvider that reads
// file from the filesystem.
func NewFileSystemProvider(inputDir string) core.FileProvider {
	return &fileSystemProvider{
		indexer: &filesindex.Indexer{
			Dir: inputDir,
		},
	}
}

func (p *fileSystemProvider) ProvideFiles(filesFilter core.Filter, processor core.Processor) error {
	for item, err := p.indexer.NextFile(); !errors.Is(err, io.EOF); item, err = p.indexer.NextFile() {
		if err != nil {
			return err
		}

		absolutePath := filepath.Join(p.indexer.Dir, item.Path())
		relativePath := item.Path()
		if filesFilter == nil || filesFilter.Accept(relativePath) {
			reader, err := open(absolutePath)
			if err != nil {
				return err
			}
			defer reader.Close()

			err = processor.ProcessFile(relativePath, reader)
			if err != nil {
				// TODO: clean output folder
				log.Printf("Error while processing file %s\n", relativePath)
				return err
			}
		}
	}

	return nil
}
