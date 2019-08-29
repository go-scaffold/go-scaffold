package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_IsTemplate_Success_ShouldCorrectlyDetectIfFileIsTemplateOrNot(t *testing.T) {
	assert.True(t, scaffold.IsTemplate("some-file.tpl"))

	assert.False(t, scaffold.IsTemplate("some-file.0tpl"))
	assert.False(t, scaffold.IsTemplate("some-file.tpl0"))
	assert.False(t, scaffold.IsTemplate("some-file.aaa"))
}
