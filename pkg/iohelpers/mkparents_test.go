package iohelpers_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_MkParents_Fail_ShouldReturnErrorIfCannotCreateDir(t *testing.T) {
	filePath := "mkparents_test.go/parent_is_a_file"

	err := iohelpers.MkParents(filePath)

	assert.NotNil(t, err)
}

func Test_MkParents_Success_ShouldCreateFolders(t *testing.T) {
	tmpDir := testutils.TempDir(t)
	outDir := tmpDir + "some/not/existing/folder/"
	filePath := filepath.Join(outDir, "some_file")
	defer os.RemoveAll(tmpDir)

	err := iohelpers.MkParents(filePath)

	assert.Nil(t, err)
	verifyDirExist(t, outDir)
}

func verifyDirExist(t *testing.T, outDir string) {
	dir, err := os.Stat(outDir)
	assert.Nil(t, err)
	assert.True(t, dir.IsDir())
}
