package scaffold_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/pasdam/go-scaffold/pkg/iohelpers"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_NewFileSystemProvider_Fail_FolderDoesNotExist(t *testing.T) {
	var filter scaffold.Filter
	processor := newMockFileProcessor()

	provider := scaffold.NewFileSystemProvider("some-non-existing-folder")
	err := provider.ProvideFiles(filter, processor)
	assert.Equal(t, "open some-non-existing-folder: no such file or directory", err.Error())
}

func TestFileSystemProvider_ProvideFiles_Fail_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter scaffold.Filter
	processor := newMockFileProcessor()
	expectedErr := errors.New("some-error")
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(expectedErr)

	provider := scaffold.NewFileSystemProvider("./testdata/file_system_provider")
	actualErr := provider.ProvideFiles(filter, processor)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, 1, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter scaffold.Filter
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider("./testdata/file_system_provider")
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, "test_folder/fileA", "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldProcessAllFileIfFilterAcceptsAll(t *testing.T) {
	filter := &mockFilter{File: "no-file-will-match"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider("./testdata/file_system_provider")
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file0", "file0-content\n")
	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, "test_folder/fileA", "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func TestFileSystemProvider_ProvideFiles_Success_ShouldNotProcessFilesIgnoredByTheFilter(t *testing.T) {
	filter := &mockFilter{File: "file0"}
	processor := newMockFileProcessor()
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

	provider := scaffold.NewFileSystemProvider("./testdata/file_system_provider")
	err := provider.ProvideFiles(filter, processor)
	assert.Nil(t, err)

	verifyProcessedFile(t, processor, "file1", "file1-content\n")
	verifyProcessedFile(t, processor, "test_folder/fileA", "fileA-content\n")
	assert.Equal(t, 0, len(processor.ReadersMap))
}

func verifyProcessedFile(t *testing.T, processor *mockFileProcessor, filePath string, content string) {
	processor.AssertCalled(t, "ProcessFile", filePath, mock.Anything)
	assert.Equal(t, content, processor.ReadersMap[filePath])
	delete(processor.ReadersMap, filePath)
}

type mockFilter struct {
	File string
}

func (m *mockFilter) Accept(filePath string) bool {
	return !strings.HasSuffix(filePath, m.File)
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

func (self *mockFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	self.ReadersMap[filePath] = iohelpers.Read(reader)
	args := self.Called(filePath, reader)
	return args.Error(0)
}
