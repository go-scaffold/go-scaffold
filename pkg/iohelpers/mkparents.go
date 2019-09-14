package iohelpers

import (
	"os"
	"path/filepath"
)

// MkParents creates intermediate directories as required.
func MkParents(filePath string) error {
	return os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
}
