package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-project-template/pkg/iohelpers"
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
	provider, err := scaffold.NewFileSystemProvider("./testdata/file_system_provider")
	assert.Nil(t, err)
	assert.True(t, provider.HasMoreFiles())

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func Test_Reset_Succeed_ResetAfterFirstRead(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./testdata/file_system_provider")

	verifyNextFile(t, provider, "file0", "file0-content\n", true)

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func Test_Reset_Succeed_ResetAfterReadSubfolder(t *testing.T) {
	provider, _ := scaffold.NewFileSystemProvider("./testdata/file_system_provider")

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
	verifyNextFile(t, provider, "file1", "file1-content\n", true)
	verifyNextFile(t, provider, "test_folder/fileA", "fileA-content\n", false)

	provider.Reset()

	verifyNextFile(t, provider, "file0", "file0-content\n", true)
}

func verifyNextFile(t *testing.T, provider scaffold.FileProvider, filePath string, content string, hasMoreFiles bool) {
	filePath, reader, err := provider.NextFile()
	assert.Nil(t, err)
	assert.Equal(t, filePath, filePath)
	assert.Equal(t, content, iohelpers.Read(reader))
}
