package filter_test

import (
	"strings"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/filter"
	"github.com/stretchr/testify/assert"
)

func Test_multiFilter_Accept_ShouldReturnsFalseIfAllFiltersDoNotAcceptTheFile(t *testing.T) {
	filter := filter.NewMultiFilter(
		&mockFilter{"file-to-exclude-0"},
		&mockFilter{"file-to-exclude-1"},
		&mockFilter{"file-to-exclude-2"},
	)
	assert.False(t, filter.Accept("file"))
}

func Test_multiFilter_Accept_ShouldReturnsTrueIfAFilterAcceptsTheFile(t *testing.T) {
	filter := filter.NewMultiFilter(
		&mockFilter{"file-to-exclude-0"},
		&mockFilter{"file-to-exclude-1"},
		&mockFilter{"file-to-exclude-2"},
	)
	assert.True(t, filter.Accept("file-to-exclude-2"))
}

type mockFilter struct {
	File string
}

func (m *mockFilter) Accept(filePath string) bool {
	return strings.HasSuffix(filePath, m.File)
}
