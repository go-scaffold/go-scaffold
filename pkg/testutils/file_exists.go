package testutils

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FileExists(t *testing.T, filePath string, expectedContent string) {
	actualContent, _ := ioutil.ReadFile(filePath)
	assert.Equal(t, expectedContent, string(actualContent))
}
