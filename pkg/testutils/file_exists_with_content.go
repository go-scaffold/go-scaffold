package testutils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FileExistsWithContent is a test helper function used to assert that the
// specified file exists and has the expected content
func FileExistsWithContent(t *testing.T, filePath string, expectedContent string) {
	actualContent, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedContent, string(actualContent))
}
