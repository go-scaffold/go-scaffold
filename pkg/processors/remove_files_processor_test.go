package processors

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/stretchr/testify/assert"
)

func TestNewRemoveFilesProcessor(t *testing.T) {
	got := NewRemoveFilesProcessor("some-root-dir").(*removeFilesProcessor)

	assert.NotNil(t, got)
	assert.Equal(t, "some-root-dir", got.rootDir)
}

func Test_removeFilesProcessor_ProcessFile_ShouldDeleteParentFolderWhenEmpty(t *testing.T) {
	path := filestest.TempFile(t, "some-test-file")
	p := NewRemoveFilesProcessor("")

	err := p.ProcessFile(path, nil)
	assert.Nil(t, err)
	filestest.PathDoesNotExist(t, filepath.Dir(path))
}

func Test_removeFilesProcessor_ProcessFile_ShouldNotDeleteParentFolderWhenNotEmpty(t *testing.T) {
	path0 := filestest.TempFile(t, "some-test-file-0")
	dir := filepath.Dir(path0)

	p := NewRemoveFilesProcessor("")

	path1 := filepath.Join(dir, "some-test-file-1")
	f, err := os.Create(path1)
	assert.Nil(t, err)
	f.Close()

	err = p.ProcessFile(path0, nil)
	assert.Nil(t, err)
	filestest.PathDoesNotExist(t, path0)
	filestest.PathExist(t, dir)
}

func Test_removeFilesProcessor_ProcessFile_ShouldUseRootDirForRelativePaths(t *testing.T) {
	path := filestest.TempFile(t, "some-test-file")
	dir := filepath.Dir(path)
	p := NewRemoveFilesProcessor(dir)

	err := p.ProcessFile("some-test-file", nil)
	assert.Nil(t, err)
	filestest.PathDoesNotExist(t, filepath.Dir(path))
}

func Test_removeFilesProcessor_ProcessFile_ShouldReturnErrorIfTheFileDoesNotExist(t *testing.T) {
	p := NewRemoveFilesProcessor("")

	err := p.ProcessFile("some-not-existing-test-file", nil)
	assert.NotNil(t, err)
	assert.Equal(t, "remove some-not-existing-test-file: no such file or directory", err.Error())
}
