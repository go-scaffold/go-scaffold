package scaffold

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/filter"
)

type fileSystemProvider struct {
	filesPath   []string
	filesInfo   []os.FileInfo
	templateDir string
}

func NewFileSystemProvider(templateDir string) FileProvider {
	return &fileSystemProvider{
		templateDir: templateDir,
	}
}

func (self *fileSystemProvider) ProvideFiles(filesFilter filter.Filter, processor FileProcessor) error {
	err := self.indexDir(self.templateDir, filesFilter)
	if err != nil {
		return err
	}

	for len(self.filesPath) > 0 {
		filePath, reader := self.nextFile(filesFilter)
		defer reader.Close()

		err = processor.ProcessFile(filePath, reader)
		if err != nil {
			// TODO: clean output folder
			return err
		}
	}
	return nil
}

func (self *fileSystemProvider) nextFile(filter filter.Filter) (string, io.ReadCloser) {
	var nextFilePath string
	var reader io.ReadCloser
	for i := 0; i < len(self.filesPath); i++ {
		nextFilePath = self.filesPath[i]

		if filter == nil || filter.Accept(nextFilePath) {
			nextFileInfo := self.filesInfo[i]

			listSize := len(self.filesPath)
			if listSize > 1 {
				self.filesPath = self.filesPath[i+1 : len(self.filesPath)]
				self.filesInfo = self.filesInfo[i+1 : len(self.filesInfo)]

			} else {
				self.filesPath = nil
				self.filesInfo = nil
			}

			if nextFileInfo.IsDir() {
				self.indexDir(nextFilePath, filter)
				return self.nextFile(filter)
			}

			reader, _ = os.Open(nextFilePath)
			nextFilePath, _ = filepath.Rel(self.templateDir, nextFilePath)
			break
		}
	}
	return nextFilePath, reader
}

func (self *fileSystemProvider) indexDir(dirPath string, filter filter.Filter) error {
	filesInfo, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	acceptedInfo := make([]os.FileInfo, 0, len(filesInfo))
	acceptedPaths := make([]string, 0, len(filesInfo))
	for i := 0; i < len(filesInfo); i++ {
		filePath := filepath.Join(dirPath, filesInfo[i].Name())
		if filter == nil || filter.Accept(filePath) {
			acceptedInfo = append(acceptedInfo, filesInfo[i])
			acceptedPaths = append(acceptedPaths, filePath)
		}
	}
	self.filesInfo = append(acceptedInfo, self.filesInfo...)
	self.filesPath = append(acceptedPaths, self.filesPath...)

	return nil
}
