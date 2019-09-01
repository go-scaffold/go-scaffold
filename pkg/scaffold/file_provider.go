package scaffold

import (
	"io"
)

type FileProvider interface {
	Reset() error
	HasMoreFiles() bool
	NextFile() (path string, reader io.ReadCloser, err error)
}
