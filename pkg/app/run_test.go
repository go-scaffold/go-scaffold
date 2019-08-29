package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-project-template/pkg/prompt"
	"github.com/pasdam/go-project-template/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

type fatalHandler struct {
	Message string
	Err     error
}

func (h *fatalHandler) Fatal(args ...interface{}) {
	h.Message = args[0].(string)
	h.Err = args[1].(error)
}

func Test_Run_Success_ValidTemplate(t *testing.T) {
	runPrompts = func(prompts []*prompt.PromptConfig) map[string]interface{} {
		data := make(map[string]interface{})
		data["text"] = "test!"
		return data
	}

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = make([]string, 5)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = filepath.Join("testdata", "valid_template")
	os.Args[3] = "--output"
	os.Args[4] = outDir

	Run()

	testutils.FileExists(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
}

func Test_Run_Fail_InvalidCliOptions(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = make([]string, 2)
	os.Args[0] = ""
	os.Args[1] = "--invalid-parameter"

	Run()

	assert.Equal(t, "Command line options error:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func Test_Run_Fail_ErrorParsingPromptFile(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = make([]string, 5)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = "invalid-folder"
	os.Args[3] = "--output"
	os.Args[4] = outDir

	Run()

	assert.Equal(t, "Unable to parse prompts.yaml file:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func Test_Run_Fail_ErrorReadingTemplateFiles(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = make([]string, 5)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = filepath.Join("testdata", "no_src")
	os.Args[3] = "--output"
	os.Args[4] = outDir

	Run()

	assert.Equal(t, "Error while creating the file provider:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func Test_Run_Fail_ErrorWhileProcessingFile(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = make([]string, 5)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = filepath.Join("testdata", "invalid_template")
	os.Args[3] = "--output"
	os.Args[4] = outDir

	Run()

	assert.Equal(t, "Error while processing files:", handler.Message)
	assert.NotNil(t, handler.Err)
}
