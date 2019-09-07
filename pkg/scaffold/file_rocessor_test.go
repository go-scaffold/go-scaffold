package scaffold_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestFileProcessor_ProcessFile_Fail_ApplyTemplateFails(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewFileProcessor(
		"invalid-data",
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.NotNil(t, err)
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "template_file.tpl"))
}

func TestFileProcessor_ProcessFile_Success_FileIsATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewFileProcessor(
		validData,
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "template_file"), "This is a *test*\n")
}

func TestFileProcessor_ProcessFile_Success_FileIsNotATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewFileProcessor(
		nil,
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("regular_file.txt", file)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "regular_file.txt"), "regular-file-content\n")
}

func TestFileProcessor_ProcessFile_Success_ShouldIgnoreFileIfItIsNotATemplateAndOnlyTemplatesIsTrue(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewFileProcessor(
		nil,
		outDir,
		&scaffold.TemplateHelper{},
		true,
	)
	err = processor.ProcessFile("regular_file.txt", file)

	assert.Nil(t, err)
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "regular_file.txt"))
}
