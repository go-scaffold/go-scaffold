package scaffold

import (
	"io"
)

// FileProcessor contains methos to process files
type FileProcessor interface {

	// ProcessFile processes the specified file, using the reader to read the content
	ProcessFile(filePath string, reader io.Reader) error
}
