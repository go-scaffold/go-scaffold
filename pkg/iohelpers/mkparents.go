package iohelpers

import (
	"os"
	"path/filepath"
)

func MkParents(filePath string) error {
	return os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
}
