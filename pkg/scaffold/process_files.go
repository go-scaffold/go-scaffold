package scaffold

func ProcessFiles(fileProvider FileProvider, config interface{}, outDir string, onlyTemplates bool) error {
	for fileProvider.HasMoreFiles() {
		filePath, reader, err := fileProvider.NextFile()
		if err != nil {
			// TODO: clean output folder
			return err
		}
		defer reader.Close()

		err = ProcessFile(reader, config, outDir, filePath, onlyTemplates)
		if err != nil {
			// TODO: clean output folder
			return err
		}
	}
	return nil
}
