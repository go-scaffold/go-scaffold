package scaffold_test

import (
	"strings"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_NewFileSystemProvider_Fail_FolderDoesNotExist(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider("some-non-existing-folder", nil)

	assert.NotNil(t, err)
	assert.Nil(t, provider)
}

func Test_NewFileSystemProvider_Succeed_FolderExist(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider(".", nil)

	assert.Nil(t, err)
	assert.NotNil(t, provider)
}

func Test_Reset_Succeed_ResetBeforeFirstRead(t *testing.T) {
	provider, err := scaffold.NewFileSystemProvider("./testdata/file_system_provider", nil)
	assert.Nil(t, err)
	assert.True(t, provider.HasMoreFiles())

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func Test_Reset_Succeed_ResetAfterFirstRead(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./testdata/file_system_provider", nil)

	verifyNextFile(t, provider, "file0", "file0-content\n", true)

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func Test_Reset_Succeed_ResetAfterReadSubfolder(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./testdata/file_system_provider", nil)

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
	verifyNextFile(t, provider, "file1", "file1-content\n", true)
	verifyNextFile(t, provider, "test_folder/fileA", "fileA-content\n", false)
	verifyNoMoreFile(t, provider)

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func Test_NextFile_Succeed_ShouldFilterFile(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./testdata/file_system_provider", &mockFilter{})

	verifyNextFile(t, provider, "file1", "file1-content\n", true)
	verifyNextFile(t, provider, "test_folder/fileA", "fileA-content\n", false)
	verifyNoMoreFile(t, provider)
}

func verifyNextFile(t *testing.T, provider scaffold.FileProvider, filePath string, content string, hasMoreFiles bool) {
	filePath, reader, err := provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, filePath, filePath)
	assert.Equal(t, content, iohelpers.Read(reader))
	assert.Equal(t, hasMoreFiles, provider.HasMoreFiles())
}

func verifyNoMoreFile(t *testing.T, provider scaffold.FileProvider) {
	filePath, reader, err := provider.NextFile()
	assert.Equal(t, "No more files", err.Error())
	assert.Empty(t, filePath)
	assert.Nil(t, reader)
	assert.Equal(t, filePath, filePath)
	assert.False(t, provider.HasMoreFiles())
}

type mockFilter struct{}

func (m *mockFilter) Accept(filePath string) bool {
	return !strings.HasSuffix(filePath, "file0")
}
