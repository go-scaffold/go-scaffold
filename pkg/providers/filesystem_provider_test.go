package providers

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/core"
	"github.com/pasdam/go-utils/pkg/assertutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFileSystemProvider_ProvideFiles_Fail_ShouldProcessAllFileIfNoFilterIsSpecified(t *testing.T) {
	var filter core.Filter
	processor := newMockFileProcessor(t)
	expectedErr := errors.New("some-error")
	processor.On("ProcessFile", mock.Anything, mock.Anything).Return(expectedErr)

	provider := NewFileSystemProvider(filepath.Join("testdata", "file_system_provider"))
	actualErr := provider.ProvideFiles(filter, processor)

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, 1, len(processor.ReadersMap))
}

func Test_fileSystemProvider_ProvideFiles(t *testing.T) {
	emptyDir := filestest.TempDir(t)

	type fields struct {
		dir string
	}
	type mocks struct {
		file         string
		filtered     bool
		processError error
		openError    error
	}
	type want struct {
		err            error
		processedFiles []string
	}
	tests := []struct {
		name   string
		fields fields
		mocks  mocks
		want   want
	}{
		{
			name: "Should return no errors if dir has no files",
			fields: fields{
				dir: emptyDir,
			},
			mocks: mocks{
				file:         "",
				filtered:     false,
				processError: nil,
				openError:    nil,
			},
			want: want{
				err:            nil,
				processedFiles: nil,
			},
		},
		{
			name: "Should propagate error if one occur while indexing folder",
			fields: fields{
				dir: "",
			},
			mocks: mocks{
				file:         "",
				filtered:     false,
				processError: nil,
				openError:    nil,
			},
			want: want{
				err:            errors.New("open : no such file or directory"),
				processedFiles: nil,
			},
		},
		{
			name: "Should propagate error if one is thrown while opening the file",
			fields: fields{
				dir: filepath.Join("testdata", "file_system_provider"),
			},
			mocks: mocks{
				file:         "",
				filtered:     false,
				processError: nil,
				openError:    errors.New("some-open-error"),
			},
			want: want{
				err:            errors.New("some-open-error"),
				processedFiles: nil,
			},
		},
		{
			name: "Should process all files in the folder and close them after use, when folder is relative to the current one",
			fields: fields{
				dir: filepath.Join("testdata", "file_system_provider"),
			},
			mocks: mocks{
				file:         "",
				filtered:     false,
				processError: nil,
				openError:    nil,
			},
			want: want{
				err:            nil,
				processedFiles: []string{"file0", "file1", filepath.Join("test_folder", "fileA")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewFileSystemProvider(tt.fields.dir)

			processor := newMockFileProcessor(t)
			if tt.mocks.processError != nil {
				processor.On("ProcessFile", tt.mocks.file, mock.Anything).Return(tt.mocks.processError)
			}
			processor.On("ProcessFile", mock.Anything, mock.Anything).Return(nil)

			mockFile, err := os.Open(filepath.Join("testdata", "file_system_provider", "file0"))
			assert.NoError(t, err)

			mockOpen(t, "", mockFile, tt.mocks.openError)

			filter := &mockFilter{
				File:    tt.mocks.file,
				Include: !tt.mocks.filtered,
			}

			err = p.ProvideFiles(filter, processor)

			assertutils.AssertEqualErrors(t, tt.want.err, err)
			processor.AssertNumberOfCalls(t, "ProcessFile", len(tt.want.processedFiles))
			for _, file := range tt.want.processedFiles {
				// mockFileCloseFn.Verify()
				processor.AssertCalled(t, "ProcessFile", file, mock.Anything)
			}
		})
	}
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
	t *testing.T

	ReadersMap map[string]string
}

func newMockFileProcessor(t *testing.T) *mockFileProcessor {
	return &mockFileProcessor{
		t:          t,
		ReadersMap: make(map[string]string),
	}
}

func (p *mockFileProcessor) ProcessFile(filePath string, reader io.Reader) error {
	content, err := io.ReadAll(reader)
	assert.NoError(p.t, err)
	p.ReadersMap[filePath] = string(content)
	args := p.Called(filePath, reader)
	return args.Error(0)
}

func mockOpen(t *testing.T, expectedName string, file *os.File, err error) {
	originalValue := open
	open = func(name string) (*os.File, error) {
		// assert.Equal(t, expectedName, name)
		return file, err
	}
	t.Cleanup(func() { open = originalValue })
}
