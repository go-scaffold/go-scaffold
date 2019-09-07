package scaffold

func ProcessFiles(fileProvider FileProvider, fileProcessor *FileProcessor) error {
	for fileProvider.HasMoreFiles() {
		filePath, reader, err := fileProvider.NextFile()
		if err != nil {
			// TODO: clean output folder
			return err
		}
		defer reader.Close()

		err = fileProcessor.ProcessFile(filePath, reader)
		if err != nil {
			// TODO: clean output folder
			return err
		}
	}
	return nil
}
