package filter_test

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/filter"
	"github.com/stretchr/testify/assert"
)

func TestNewPatternFilter_Fail_ShouldReturnErrorIfAPatternIsInvalid(t *testing.T) {
	tests := []struct {
		inclusive bool
	}{
		{inclusive: false},
		{inclusive: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			filter, err := filter.NewPatternFilter(tt.inclusive, "pattern1", "invalid pattern [")
			assert.Equal(t, "error parsing regexp: missing closing ]: `[`", err.Error())
			assert.Nil(t, filter)
		})
	}
}

func Test_PatternFilter_Accept_ShouldReturnFalseIfItIsExclusiveAndTheFileNameMatchesThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter(false, "expression-to-exclude-1", ".*file-to-match.*")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.False(t, filter.Accept("file-to-match"))
}

func Test_PatternFilter_Accept_ShouldReturnTrueIfItIsExclusiveAndTheFileNameDoesNotMatcheThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter(false, "expression-to-exclude-1", "expression-to-exclude-2")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.True(t, filter.Accept("file-to-match"))
}

func Test_PatternFilter_Accept_ShouldReturnTrueIfItIsInclusiveAndTheFileNameMatchesThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter(true, "expression-to-match-1", ".*file-to-match.*")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.True(t, filter.Accept("file-to-match"))
}

func Test_PatternFilter_Accept_ShouldReturnFalseIfItIsInclusiveAndTheFileNameDoesNotMatcheThePatterns(t *testing.T) {
	filter, err := filter.NewPatternFilter(true, "expression-to-match-1", "expression-to-match-2")
	assert.Nil(t, err)
	assert.NotNil(t, filter)

	assert.False(t, filter.Accept("file-to-match"))
}

func Test_PatternFilter_NewInstance_ShouldOverwriteConfig(t *testing.T) {
	tests := []struct {
		name         string
		inclusive    bool
		shouldAccept bool
		fileName     string
	}{
		{"0", false, true, "not-matching"},
		{"1", false, false, "expression-to-match-1"},
		{"2", true, false, "not-matching"},
		{"3", true, true, "expression-to-match-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := filter.NewPatternFilter(false, "expression-to-match-1", "expression-to-match-2")
			assert.Nil(t, err)
			assert.NotNil(t, f)

			copyFilter := filter.NewPatternFilterFromInstance(f, tt.inclusive)

			assert.Equal(t, tt.shouldAccept, copyFilter.Accept(tt.fileName))
		})
	}
}
