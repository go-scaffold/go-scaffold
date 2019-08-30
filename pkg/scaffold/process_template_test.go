package scaffold_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_ProcessTemplate_fail_shouldReturnErrorIfApplyingTheTemplateFailed(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	reader, err := scaffold.ProcessTemplate(file, "invalid-data")

	assert.NotNil(t, err)
	assert.Nil(t, reader)
}

func Test_ProcessTemplate_success_shouldCreateAReaderForTheGeneratedContent(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	reader, err := scaffold.ProcessTemplate(file, validData)

	assert.Nil(t, err)
	readContent, err := ioutil.ReadAll(reader)
	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*\n", string(readContent))
}
