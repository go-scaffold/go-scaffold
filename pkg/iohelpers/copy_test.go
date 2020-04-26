package iohelpers_test

import (
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_Copy_Fail_ShouldReturnErrorIfFileDoesNotExist(t *testing.T) {
	outDir := testutils.TempDir(t)

	filePath := "non-exsisting-file"
	err := iohelpers.Copy(filePath, filepath.Join(outDir, filePath))

	assert.NotNil(t, err)
}

func Test_Copy_Success_ShouldCopyExistingFile(t *testing.T) {
	outDir := testutils.TempDir(t)

	filePath := filepath.Join("testdata", "file_to_read.txt")
	outFile := filepath.Join(outDir, filePath)
	err := iohelpers.Copy(filePath, outFile)

	assert.Nil(t, err)
	testutils.FileExistsWithContent(t, outFile, "file-to-read-content\n")
}
