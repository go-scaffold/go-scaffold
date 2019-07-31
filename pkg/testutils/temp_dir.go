package testutils

import (
	"crypto/rand"
	"fmt"
	"os"
	"testing"
)

func TempDir(t *testing.T) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		t.FailNow()
	}

	dirPath := fmt.Sprintf("/var/tmp/tempdir-%X/", b[0:])
	os.MkdirAll(dirPath, os.ModePerm)

	return dirPath
}
