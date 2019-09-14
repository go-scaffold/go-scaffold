package scaffold

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/filter"
)

type fileSystemProvider struct {
	cleanFilter filter.Filter
	filesPath   []string
	filesInfo   []os.FileInfo
	inputDir    string
}

// NewFileSystemProvider creates a new instance of a FileProvider that reads file from the filesystem.
// If cleanFilter is specified it will be used to remove all the inpout files that it matches.
func NewFileSystemProvider(inputDir string, cleanFilter filter.Filter) FileProvider {
	return &fileSystemProvider{
		cleanFilter: cleanFilter,
		inputDir:    inputDir,
	}
}

func (p *fileSystemProvider) ProvideFiles(filesFilter filter.Filter, processor FileProcessor) error {
	err := p.indexDir(p.inputDir)
	if err != nil {
		return err
	}

	for len(p.filesPath) > 0 {
		filePath, reader := p.nextFile()
		if reader != nil {
			defer reader.Close()
		}

		relativePath, _ := filepath.Rel(p.inputDir, filePath)
		if reader != nil && (filesFilter == nil || filesFilter.Accept(relativePath)) {
			err = processor.ProcessFile(relativePath, reader)
			if err != nil {
				// TODO: clean output folder
				log.Printf("Error while processing file %s\n", relativePath)
				return err
			}
		}

		if p.cleanFilter != nil && p.cleanFilter.Accept(filePath) {
			os.RemoveAll(filePath)
		}
	}
	return nil
}

func (p *fileSystemProvider) nextFile() (string, io.ReadCloser) {
	nextFileInfo := p.filesInfo[0]
	nextFilePath := p.filesPath[0]

	listSize := len(p.filesPath)
	if listSize > 1 {
		p.filesPath = p.filesPath[1:len(p.filesPath)]
		p.filesInfo = p.filesInfo[1:len(p.filesInfo)]

	} else {
		p.filesPath = nil
		p.filesInfo = nil
	}

	var reader io.ReadCloser

	if nextFileInfo.IsDir() {
		p.indexDir(nextFilePath)
		return nextFilePath, nil
	}

	reader, _ = os.Open(nextFilePath)
	return nextFilePath, reader
}

func (p *fileSystemProvider) indexDir(dirPath string) error {
	filesInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	acceptedInfo := make([]os.FileInfo, 0, len(filesInfo))
	acceptedPaths := make([]string, 0, len(filesInfo))
	for i := 0; i < len(filesInfo); i++ {
		filePath := filepath.Join(dirPath, filesInfo[i].Name())
		acceptedInfo = append(acceptedInfo, filesInfo[i])
		acceptedPaths = append(acceptedPaths, filePath)
	}
	p.filesInfo = append(acceptedInfo, p.filesInfo...)
	p.filesPath = append(acceptedPaths, p.filesPath...)

	return nil
}
