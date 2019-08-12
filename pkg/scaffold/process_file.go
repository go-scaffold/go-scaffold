package scaffold

import (
	"path/filepath"

	"github.com/pasdam/go-project-template/pkg/iohelpers"
)

func ProcessFile(reader FileReader, config interface{}, outDir string, filePath string) error {
	if IsTemplate(filePath) {

		content, err := ApplyTemplate(iohelpers.Read(reader), config)
		if err != nil {
			return err
		}

		err = writeTextFile(content, filepath.Join(outDir, OutputFilePath(filePath)))
		if err != nil {
			return err
		}

	} else {
		return iohelpers.WriteFile(reader, filepath.Join(outDir, filePath))
	}

	return nil
}
