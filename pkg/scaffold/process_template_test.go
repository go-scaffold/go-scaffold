package scaffold

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessTemplate_fail_shouldReturnErrorIfApplyingTheTemplateFailed(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	reader, err := processTemplate(file, "invalid-data")

	assert.NotNil(t, err)
	assert.Nil(t, reader)
}

func Test_ProcessTemplate_success_shouldCreateAReaderForTheGeneratedContent(t *testing.T) {
	file, err := os.Open("testdata/template_file.tpl")
	assert.Nil(t, err)
	defer file.Close()

	reader, err := processTemplate(file, struct{ Text string }{Text: "*test*"})

	assert.Nil(t, err)
	readContent, err := ioutil.ReadAll(reader)
	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*\n", string(readContent))
}
