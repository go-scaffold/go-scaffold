package scaffold

import (
	"io"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

func ProcessFile(reader io.Reader, config interface{}, outDir string, filePath string, onlyTemplate bool) error {
	var err error
	if IsTemplate(filePath) {
		reader, err = ProcessTemplate(reader, config)
		if err != nil {
			return err
		}
		filePath = OutputFilePath(filePath)

	} else if onlyTemplate {
		return nil
	}

	return iohelpers.WriteFile(reader, filepath.Join(outDir, filePath))
}
