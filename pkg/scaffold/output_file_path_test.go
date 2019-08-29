package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_OutputFilePath_Success_ShouldTrimSuffix(t *testing.T) {
	assert.Equal(t, "some-file.txt", scaffold.OutputFilePath("some-file.txt.tpl"))
	assert.Equal(t, "some-other-file.txt", scaffold.OutputFilePath("some-other-file.txt"))
}
