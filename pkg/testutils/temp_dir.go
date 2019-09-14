package testutils

import (
	"io/ioutil"
	"os"
	"testing"
)

// TempDir is a test helper function used to create a temporary dir
func TempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-scaffold")
	if err != nil {
		t.FailNow()
	}

	return dir
}
