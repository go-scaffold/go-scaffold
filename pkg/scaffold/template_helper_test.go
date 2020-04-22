package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_Accept_Success_ShouldCorrectlyDetectIfFileIsTemplateOrNot(t *testing.T) {
	helper := scaffold.TemplateHelper{}

	assert.True(t, helper.Accept("some-file.tpl"))

	assert.False(t, helper.Accept("some-file.0tpl"))
	assert.False(t, helper.Accept("some-file.tpl0"))
	assert.False(t, helper.Accept("some-file.aaa"))
}

func Test_OutputFilePath_Success_ShouldTrimSuffix(t *testing.T) {
	helper := scaffold.TemplateHelper{}

	assert.Equal(t, "some-file.txt", helper.OutputFilePath("some-file.txt.tpl"))
	assert.Equal(t, "some-other-file.txt", helper.OutputFilePath("some-other-file.txt"))
}
