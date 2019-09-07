package scaffold

type FileProvider interface {
	ProvideFiles(filesFilter Filter, processor FileProcessor) error
}
