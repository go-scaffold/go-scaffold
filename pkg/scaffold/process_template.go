package scaffold

import (
	"io"
	"strings"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

func ProcessTemplate(reader io.Reader, config interface{}) (io.Reader, error) {
	content, err := ApplyTemplate(iohelpers.Read(reader), config)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(content), nil
}
