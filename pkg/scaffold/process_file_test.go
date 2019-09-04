package scaffold_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
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
		false,
	)

	assert.NotNil(t, err)
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "template_file.tpl"))
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
		false,
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
		false,
	)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "testdata/regular_file.txt"), "regular-file-content\n")
}

func Test_ProcessFile_Success_ShouldIgnoreFileIfItIsNotATemplateAndOnlyTemplatesIsTrue(t *testing.T) {
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
		true,
	)

	assert.Nil(t, err)
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "testdata/regular_file.txt"))
}
