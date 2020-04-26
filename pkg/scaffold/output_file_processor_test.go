package scaffold_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestOutputFileProcessor_ProcessFile_Fail_ApplyTemplateFails(t *testing.T) {
	outDir := testutils.TempDir(t)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		"invalid-data",
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.NotNil(t, err)
	testutils.PathDoesNotExist(t, filepath.Join(outDir, "template_file.tpl"))
}

func TestOutputFileProcessor_ProcessFile_Success_FileIsATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)

	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		struct{ Text string }{Text: "*test*"},
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("template_file.tpl", file)

	assert.Nil(t, err)
	testutils.FileExistsWithContent(t, filepath.Join(outDir, "template_file"), "This is a *test*\n")
}

func TestOutputFileProcessor_ProcessFile_Success_FileIsNotATemplate(t *testing.T) {
	outDir := testutils.TempDir(t)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		nil,
		outDir,
		&scaffold.TemplateHelper{},
		false,
	)
	err = processor.ProcessFile("regular_file.txt", file)

	assert.Nil(t, err)
	testutils.FileExistsWithContent(t, filepath.Join(outDir, "regular_file.txt"), "regular-file-content\n")
}

func TestOutputFileProcessor_ProcessFile_Success_ShouldIgnoreFileIfItIsNotATemplateAndOnlyTemplatesIsTrue(t *testing.T) {
	outDir := testutils.TempDir(t)

	file, err := os.Open("testdata/regular_file.txt")
	assert.Nil(t, err)
	defer file.Close()

	processor := scaffold.NewOutputFileProcessor(
		nil,
		outDir,
		&scaffold.TemplateHelper{},
		true,
	)
	err = processor.ProcessFile("regular_file.txt", file)

	assert.Nil(t, err)
	testutils.PathDoesNotExist(t, filepath.Join(outDir, "regular_file.txt"))
}
