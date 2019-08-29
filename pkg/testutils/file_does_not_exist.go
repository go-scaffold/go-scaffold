package testutils

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FileDoesNotExist(t *testing.T, filePath string) {
	_, err := os.Stat(filePath)
	assert.NotNil(t, err)
	assert.True(t, strings.HasSuffix(err.Error(), ": no such file or directory"))
}
