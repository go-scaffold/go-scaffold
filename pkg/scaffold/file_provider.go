package scaffold

type FileProvider interface {
	Reset() error
	HasMoreFiles() bool
	NextFile() (path string, reader FileReader, err error)
}
