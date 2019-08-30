package iohelpers_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_Copy_Fail_ShouldReturnErrorIfFileDoesNotExist(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	filePath := "non-exsisting-file"
	err := iohelpers.Copy(filePath, filepath.Join(outDir, filePath))

	assert.NotNil(t, err)
}

func Test_Copy_Success_ShouldCopyExistingFile(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	filePath := filepath.Join("testdata", "file_to_read.txt")
	outFile := filepath.Join(outDir, filePath)
	err := iohelpers.Copy(filePath, outFile)

	assert.Nil(t, err)
	testutils.FileExists(t, outFile, "file-to-read-content\n")
}
