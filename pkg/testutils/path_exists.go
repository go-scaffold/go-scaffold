package testutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// PathExist is a test helper function used to assert that the specified path
// exists
func PathExist(t *testing.T, filePath string) {
	info, err := os.Stat(filePath)
	assert.Nil(t, err)
	assert.NotNil(t, info)
}
