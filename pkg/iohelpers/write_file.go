package iohelpers

import (
	"io"
	"log"
	"os"
)

// WriteFile writes the content of Reader to the specified destination file
func WriteFile(reader io.Reader, dst string) error {
	err := MkParents(dst)
	if err != nil {
		log.Printf("Error while creating parents folder of %s\n", dst)
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
