package scaffold

import (
	"github.com/pasdam/go-project-template/pkg/iohelpers"
)

func ProcessFile(reader FileReader, config interface{}, outDir string, filePath string) error {
	if IsTemplate(filePath) {

		content, err := ApplyTemplate(iohelpers.Read(reader), config)
		if err != nil {
			return err
		}

		err = writeTextFile(content, outDir+OutputFilePath(filePath))
		if err != nil {
			return err
		}

	} else {
		return iohelpers.WriteFile(reader, outDir+filePath)
	}

	return nil
}
