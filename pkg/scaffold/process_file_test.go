package scaffold_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-project-template/pkg/scaffold"
	"github.com/pasdam/go-project-template/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessFile_Fail_ApplyTemplateFails(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		"invalid-data",
		outDir,
		"template_file.tpl",
	)

	assert.NotNil(t, err)
	verifyOutputFileDoesNotExist(t, outDir, "template_file.tpl")
}

func Test_ProcessFile_Fail_CantWriteOutputFile(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		validData,
		outDir,
		"some-non-existent-folder/testdata/template_file.tpl",
	)

	assert.NotNil(t, err)
	verifyOutputFileDoesNotExist(t, outDir, "testdata/template_file.tpl")
}

func Test_ProcessFile_Success_FileIsATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		validData,
		outDir,
		"template_file.tpl",
	)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "template_file"), "This is a *test*\n")
}

func Test_ProcessFile_Success_FileIsNotATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	err = scaffold.ProcessFile(
		file,
		nil,
		outDir,
		"testdata/regular_file.txt",
	)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "testdata/regular_file.txt"), "regular-file-content\n")
}

func verifyOutputFileDoesNotExist(t *testing.T, outDir string, filePath string) {
	_, err := os.Stat(filepath.Join(outDir, filePath))
	assert.NotNil(t, err)
}
