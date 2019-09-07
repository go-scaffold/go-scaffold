package scaffold

import (
	"io"
)

type FileProcessor interface {
	ProcessFile(filePath string, reader io.Reader) error
}
