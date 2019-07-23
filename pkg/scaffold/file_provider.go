package scaffold

type FileProvider interface {
	Reset() error
	HasMoreFiles() bool
	NextFile() (string, error)
}
