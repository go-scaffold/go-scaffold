package scaffold

func ProcessFiles(fileProvider FileProvider, config interface{}, outDir string, onlyTemplates bool, templateHelper *TemplateHelper) error {
	for fileProvider.HasMoreFiles() {
		filePath, reader, err := fileProvider.NextFile()
		if err != nil {
			// TODO: clean output folder
			return err
		}
		defer reader.Close()

		err = ProcessFile(reader, config, outDir, filePath, onlyTemplates, templateHelper)
		if err != nil {
			// TODO: clean output folder
			return err
		}
	}
	return nil
}
