package templates

import (
	"io"
	"log"
	"strings"

	"github.com/pasdam/go-io-utilx/pkg/ioutilx"
)

// ProcessTemplate processes the template using the specified data
func ProcessTemplate(reader io.Reader, data interface{}) (io.Reader, error) {
	content, err := applyTemplate(ioutilx.ReaderToString(reader), data)
	if err != nil {
		log.Println("Error while generating output file from template")
		return nil, err
	}

	return strings.NewReader(content), nil
}
