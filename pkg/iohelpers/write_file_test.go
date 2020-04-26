package iohelpers_test

import (
	"os"
	"strings"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_WriteFile_Fail_ShouldReturnErrorIfItCannotCreateParents(t *testing.T) {
	dstPath := "mkparents_test.go/parent_is_a_file"
	reader := strings.NewReader("reader-string")

	err := iohelpers.WriteFile(reader, dstPath)

	assert.NotNil(t, err)
}

func Test_WriteFile_Fail_ShouldReturnErrorIfItCannotCreateFile(t *testing.T) {
	tmpDir := testutils.TempDir(t)
	defer os.RemoveAll(tmpDir)
	dstPath := tmpDir + ""
	expectedContent := "reader-string"
	reader := strings.NewReader(expectedContent)

	err := iohelpers.WriteFile(reader, dstPath)

	assert.NotNil(t, err)
}

func Test_WriteFile_Success_ShouldCopyFile(t *testing.T) {
	tmpDir := testutils.TempDir(t)
	defer os.RemoveAll(tmpDir)
	dstPath := tmpDir + "out_file"
	expectedContent := "reader-string"
	reader := strings.NewReader(expectedContent)

	err := iohelpers.WriteFile(reader, dstPath)

	assert.Nil(t, err)
	testutils.FileExistsWithContent(t, dstPath, expectedContent)
}
