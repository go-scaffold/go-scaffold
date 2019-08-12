package testutils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(os.TempDir(), "go-scaffold")
	if err != nil {
		t.FailNow()
	}

	return dir
}
