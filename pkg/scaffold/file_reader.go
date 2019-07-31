package scaffold

import (
	"io"
)

type FileReader interface {
	io.Reader
	io.Closer
}
