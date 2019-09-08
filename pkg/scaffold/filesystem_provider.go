package scaffold

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/filter"
)

type fileSystemProvider struct {
	cleanFilter filter.Filter
	filesPath   []string
	filesInfo   []os.FileInfo
	templateDir string
}

func NewFileSystemProvider(templateDir string, cleanFilter filter.Filter) FileProvider {
	return &fileSystemProvider{
		cleanFilter: cleanFilter,
		templateDir: templateDir,
	}
}

func (self *fileSystemProvider) ProvideFiles(filesFilter filter.Filter, processor FileProcessor) error {
	err := self.indexDir(self.templateDir)
	if err != nil {
		return err
	}

	for len(self.filesPath) > 0 {
		filePath, reader := self.nextFile()
		if reader != nil {
			defer reader.Close()
		}

		relativePath, _ := filepath.Rel(self.templateDir, filePath)
		if reader != nil && (filesFilter == nil || filesFilter.Accept(relativePath)) {
			err = processor.ProcessFile(relativePath, reader)
			if err != nil {
				// TODO: clean output folder
				return err
			}
		}

		if self.cleanFilter != nil && self.cleanFilter.Accept(filePath) {
			os.RemoveAll(filePath)
		}
	}
	return nil
}

func (self *fileSystemProvider) nextFile() (string, io.ReadCloser) {
	nextFileInfo := self.filesInfo[0]
	nextFilePath := self.filesPath[0]

	listSize := len(self.filesPath)
	if listSize > 1 {
		self.filesPath = self.filesPath[1:len(self.filesPath)]
		self.filesInfo = self.filesInfo[1:len(self.filesInfo)]

	} else {
		self.filesPath = nil
		self.filesInfo = nil
	}

	var reader io.ReadCloser

	if nextFileInfo.IsDir() {
		self.indexDir(nextFilePath)
		return nextFilePath, nil
	}

	reader, _ = os.Open(nextFilePath)
	return nextFilePath, reader
}

func (self *fileSystemProvider) indexDir(dirPath string) error {
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
	self.filesInfo = append(acceptedInfo, self.filesInfo...)
	self.filesPath = append(acceptedPaths, self.filesPath...)

	return nil
}
