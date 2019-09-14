package iohelpers

import (
	"os"
)

// Copy copies the source file to the specified destination
func Copy(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	return WriteFile(file, destination)
}
