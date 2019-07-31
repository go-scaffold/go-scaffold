package iohelpers

import (
	"os"
	"strings"
)

func MkParents(filePath string) error {
	// TODO: improve this
	parts := strings.Split(filePath, "/")
	parts = parts[:len(parts)-1]
	filePath = strings.Join(parts, "/")

	return os.MkdirAll(filePath, os.ModePerm)
}
