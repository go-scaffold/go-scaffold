package scaffold

import (
	"strings"
)

func ProcessFiles(fileProvider FileProvider, config interface{}, outDir string) error {
	if !strings.HasSuffix(outDir, "/") {
		outDir = outDir + "/"
	}

	for fileProvider.HasMoreFiles() {
		filePath, reader, err := fileProvider.NextFile()
		if err != nil {
			// TODO: clean output folder
			return err
		}
		defer reader.Close()

		err = ProcessFile(reader, config, outDir, filePath)
		if err != nil {
			// TODO: clean output folder
			return err
		}
	}
	return nil
}
