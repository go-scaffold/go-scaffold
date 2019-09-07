package filter_test

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/filter"
	"github.com/stretchr/testify/assert"
)

func Test_NewPatternFilter_Success_ValidPatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter("pattern1", "pattern1")

	assert.Nil(t, err)
	assert.NotNil(t, filter)
}

func Test_NewPatternFilter_Fail_ShouldReturnErrorIfOneOFThePatternIsInvalid(t *testing.T) {
	filter, err := filter.NewPatternFilter("pattern1", "invalid pattern [")

	assert.NotNil(t, err)
	assert.Nil(t, filter)
}

func Test_Accept_ShouldReturnFalseIfTheFileNameMatchesThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter("non-matching-expression", ".*file-to-match.*")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.False(t, filter.Accept("file-to-match"))
}

func Test_Accept_ShouldReturnTrueIfTheFileNameDoesNotMatcheThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter("non-matching-expression-1", "non-matching-expression-2")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.True(t, filter.Accept("file-to-match"))
}
