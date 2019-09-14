package scaffold

import (
	"io"
	"log"
	"strings"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
)

func processTemplate(reader io.Reader, config interface{}) (io.Reader, error) {
	content, err := applyTemplate(iohelpers.Read(reader), config)
	if err != nil {
		log.Println("Error while generating output file from template")
		return nil, err
	}

	return strings.NewReader(content), nil
}
