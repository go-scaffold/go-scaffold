package iohelpers

import (
	"io/ioutil"
	"os"
)

// RemoveDirIfEmpty removes the specified directory if it's empty. Note that it
// will not return error if the directory is not empty, it will simply return
// nil, same thing if id doesn't exist already.
func RemoveDirIfEmpty(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return nil
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return os.Remove(path)
	}

	return nil
}
