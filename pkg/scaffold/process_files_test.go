package scaffold_test

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_ProcessFiles_Fail_ShouldReturnErrorIfFileProviderReturnsError(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)
	expectedError := errors.New("Expected error")
	provider := &mockFileProvider{err: expectedError}

	err := scaffold.ProcessFiles(provider, validData, outDir, false)

	assert.Equal(t, expectedError, err)
}

func Test_ProcessFiles_Fail_ShouldReturnErrorIfCannotProcessFile(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)
	provider := &mockFileProvider{}

	err := scaffold.ProcessFiles(provider, validData, outDir, false)

	assert.NotNil(t, err)
}

func Test_ProcessFiles_Success_ShouldCreateOutputFiles_DirWithoutSuffix(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	provider, _ := scaffold.NewFileSystemProvider("testdata/")

	err := scaffold.ProcessFiles(provider, validData, outDir, false)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "regular_file.txt"), "regular-file-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "template_file"), "This is a *test*\n")
	testutils.FileExists(t, filepath.Join(outDir, "file_system_provider/file0"), "file0-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "file_system_provider/file1"), "file1-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "file_system_provider/test_folder/fileA"), "fileA-content\n")
}

func Test_ProcessFiles_Success_ShouldCreateOutputFiles_DirWithoutSuffixOnlyTemplates(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	provider, _ := scaffold.NewFileSystemProvider("testdata/")

	err := scaffold.ProcessFiles(provider, validData, outDir, true)

	assert.Nil(t, err)

	testutils.FileDoesNotExist(t, filepath.Join(outDir, "regular_file.txt"))
	testutils.FileExists(t, filepath.Join(outDir, "template_file"), "This is a *test*\n")
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "file_system_provider/file0"))
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "file_system_provider/file1"))
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "file_system_provider/test_folder/fileA"))
}

func Test_ProcessFiles_Success_ShouldCreateTheOutputFiles_DirWithSuffix(t *testing.T) {
	outDir := testutils.TempDir(t) + "/"
	defer os.RemoveAll(outDir)

	provider, _ := scaffold.NewFileSystemProvider("testdata/")

	err := scaffold.ProcessFiles(provider, validData, outDir, false)

	assert.Nil(t, err)
	testutils.FileExists(t, filepath.Join(outDir, "/regular_file.txt"), "regular-file-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "/template_file"), "This is a *test*\n")
	testutils.FileExists(t, filepath.Join(outDir, "/file_system_provider/file0"), "file0-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "/file_system_provider/file1"), "file1-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "/file_system_provider/test_folder/fileA"), "fileA-content\n")
}

type mockFileProvider struct {
	mock.Mock

	err error
}

func (m *mockFileProvider) HasMoreFiles() bool {
	return true
}

func (m *mockFileProvider) NextFile() (path string, reader io.ReadCloser, err error) {
	return "", &mockReader{}, m.err
}

func (m *mockFileProvider) Reset() error {
	return nil
}

type mockReader struct{}

func (m *mockReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m *mockReader) Close() error {
	return nil
}
