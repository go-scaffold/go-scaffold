package iohelpers

import (
	"os"
)

func Copy(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}

	return WriteFile(file, destination)
}
