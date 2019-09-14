package scaffold_test

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/otiai10/copy"
	"github.com/pasdam/go-scaffold/pkg/filter"
	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_NewFileSystemProvider_Fail_FolderDoesNotExist(t *testing.T) {
	var filter filter.Filter
	processor := newMockFileProcessor()

	provider := scaffold.NewFileSystemProvider("some-non-existing-folder", nil)
	err := provider.ProvideFiles(filter, processor)
	assert.Equal(t, "open some-non-existing-folder: no such file or directory", err.Error())
}

func TestFileSystemProvider_ProvideFiles_Fail_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter filter.Filter
	processor := newMockFileProcessor()
	expectedErr := errors.New("some-error")
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(expectedErr)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"), nil)
	actualErr := provider.ProvideFiles(filter, processor)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, 1, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter filter.Filter
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"), nil)
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfFilterAcceptsAll(t *testing.T) {
	filter := &mockFilter{File: "no-file-will-match"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"), nil)
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldNotProcessFilesIgnoredByTheFilter(t *testing.T) {
	filter := &mockFilter{File: "file0"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"), nil)
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldCleanSourceFiles(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	copy.Copy(filepath.Join("testdata", "file_system_provider"), outDir)
	testutils.FileExists(t, filepath.Join(outDir, "file0"), "file0-content\n")

	fileToCleanFilter := &mockFilter{
		File:    "file0",
		Include: true,
	}
	fileToIgnoreFilter := &mockFilter{
		File:    "file0",
		Include: false,
	}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(outDir, fileToCleanFilter)
	err := provider.ProvideFiles(fileToIgnoreFilter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, filepath.Join("test_folder", "fileA"), "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))

	testutils.FileExists(t, filepath.Join(outDir, "file1"), "file1-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "test_folder", "fileA"), "fileA-content\n")
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "file0"))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldCleanSourceFolder(t *testing.T) {
	outDir := testutils.TempDir(t)
	defer os.RemoveAll(outDir)

	copy.Copy(filepath.Join("testdata", "file_system_provider"), outDir)
	testutils.FileExists(t, filepath.Join(outDir, "file0"), "file0-content\n")

	fileToCleanFilter := &mockFilter{
		File:    "test_folder",
		Include: true,
	}
	fileToIgnoreFilter := &mockFilter{
		File:    "test_folder",
		Include: false,
	}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider(outDir, fileToCleanFilter)
	err := provider.ProvideFiles(fileToIgnoreFilter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))

	testutils.FileExists(t, filepath.Join(outDir, "file0"), "file0-content\n")
	testutils.FileExists(t, filepath.Join(outDir, "file1"), "file1-content\n")
	testutils.FileDoesNotExist(t, filepath.Join(outDir, "test_folder"))
}

func verifyProcessedFile(t *testing.T, processor *mockFileProcessor, filePath string, content string) {
	processor.AssertCalled(t, "ProcessFile", filePath, mock.Anything)
	assert.Equal(t, content, processor.ReadersMap[filePath])
	delete(processor.ReadersMap, filePath)
}

type mockFilter struct {
	File    string
	Include bool
}

func (m *mockFilter) Accept(filePath string) bool {
	if m.Include {
		return strings.Contains(filePath, m.File)

	} else {
		return !strings.Contains(filePath, m.File)
	}
}

type mockFileProcessor struct {
	mock.Mock

	ReadersMap map[string]string
}

func newMockFileProcessor() *mockFileProcessor {
	return &mockFileProcessor{
		ReadersMap: make(map[string]string),
	}
}

func (p *mockFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	p.ReadersMap[filePath] = iohelpers.Read(reader)
	args := p.Called(filePath, reader)
	return args.Error(0)
}
