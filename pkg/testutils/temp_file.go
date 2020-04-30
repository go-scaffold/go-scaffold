package testutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TempFile is a test helper function used to create a temporary file
func TempFile(t *testing.T, name string) string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-scaffold")
	if err != nil {
		t.Error("Unable to create temp dir", err)
		return ""
	}

	t.Cleanup(func() {
		defer os.RemoveAll(dir)
	})

	path := filepath.Join(dir, name)

	err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		t.Error("Unable to create temp dir", err)
		return ""
	}

	file, err := os.Create(path)
	if err != nil {
		t.Error("Unable to create temp file", err)
		return ""
	}
	file.Close()

	return path
}
