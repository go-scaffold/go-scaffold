package iohelpers

import (
	"io"
	"os"
)

func WriteFile(reader io.Reader, dst string) error {
	err := MkParents(dst)
	if err != nil {
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, reader)
	return err
}
