package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/prompt"
	"github.com/pasdam/go-scaffold/pkg/testutils"
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
	mockPrompt()

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := mockArguments(filepath.Join("testdata", "valid_template"), outDir, false)
	defer func() { os.Args = oldArgs }()

	Run()

	testutils.FileExists(t, filepath.Join(outDir, "file.txt"), "This is a test!\n")
	testutils.FileExists(t, filepath.Join(outDir, "normal_file.txt"), "normal-file-content\n")
	testutils.FileDoesNotExist(t, filepath.Join(outDir, ".go-scaffold"))
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

	oldArgs := mockArguments(filepath.Join("testdata", "valid_template", ".go-scaffold"), outDir, false)
	defer func() { os.Args = oldArgs }()

	Run()

	assert.Equal(t, "Unable to parse prompts.yaml file:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func Test_Run_Fail_NotExistingFolder(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := mockArguments(filepath.Join("testdata", "invalid-folder"), outDir, false)
	defer func() { os.Args = oldArgs }()

	Run()

	assert.Equal(t, "Unable to parse prompts.yaml file:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func Test_Run_Fail_ErrorWhileProcessingFiles(t *testing.T) {
	handler := &fatalHandler{}
	fatal = handler.Fatal

	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	oldArgs := mockArguments(filepath.Join("testdata", "invalid_template"), outDir, false)
	defer func() { os.Args = oldArgs }()

	Run()

	assert.Equal(t, "Error while processing files:", handler.Message)
	assert.NotNil(t, handler.Err)
}

func mockPrompt() {
	runPrompts = func(prompts []*prompt.PromptConfig) map[string]interface{} {
		data := make(map[string]interface{})
		data["text"] = "test!"
		return data
	}
}

func mockArguments(templateDir string, outDir string, withRemoveSource bool) []string {
	oldArgs := os.Args

	os.Args = make([]string, 7)
	os.Args[0] = ""
	os.Args[1] = "--template"
	os.Args[2] = templateDir
	os.Args[3] = "--output"
	os.Args[4] = outDir
	if withRemoveSource {
		os.Args[5] = "--remove-source"
		os.Args[6] = outDir
	}

	return oldArgs
}
