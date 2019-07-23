package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-project-template/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_NewFileSystemProvider_Fail_FolderDoesNotExist(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider("some-non-existing-folder")

	assert.NotNil(t, err)
	assert.Nil(t, provider)
}

func Test_NewFileSystemProvider_Succeed_FolderExist(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider(".")

	assert.Nil(t, err)
	assert.NotNil(t, provider)
}

func Test_Reset_Succeed_ResetBeforeFirstRead(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider("./test")
	assert.Nil(t, err)
	assert.True(t, provider.HasMoreFiles())

	provider.Reset()

	file, err := provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, "file0", file)
}

func Test_Reset_Succeed_ResetAfterFirstRead(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./test")

	file, err := provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, "file0", file)

	provider.Reset()

	file, err = provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, "file0", file)
}

func Test_Reset_Succeed_ResetAfterReadSubfolder(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./test")

	file, err := provider.NextFile()
	assert.Equal(t, "file0", file)
	file, err = provider.NextFile()
	assert.Equal(t, "file1", file)

	file, err = provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, "test_folder/fileA", file)

	provider.Reset()

	file, err = provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, "file0", file)
}
