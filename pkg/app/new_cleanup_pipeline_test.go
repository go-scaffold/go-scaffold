package app

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/mockit/matchers/argument"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_newCleanupPipeline_ShouldReturnErrorIfOneOccursWhenCreatingTheFilter(t *testing.T) {
	expected := errors.New("some-filter-error")
	mockit.MockFunc(t, filters.NewPatternFilter).With(true, argument.Any).Return(nil, expected)

	got, err := newCleanupPipeline("some-source-dir")

	assert.Nil(t, got)
	assert.Equal(t, expected, err)
}

func Test_newCleanupPipeline_ShouldDeleteGoScaffoldFiles(t *testing.T) {
	relPath := filepath.Join(".go-scaffold", "some-file")
	path := filestest.TempFile(t, relPath)
	filestest.PathExist(t, path)
	dir := filepath.Dir(filepath.Dir(path))

	got, err := newCleanupPipeline(dir)

	assert.NotNil(t, got)
	assert.Nil(t, err)

	err = got.ProcessFile(relPath, nil)

	assert.Nil(t, err)
	filestest.PathDoesNotExist(t, path)
}
