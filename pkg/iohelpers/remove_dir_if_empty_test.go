package iohelpers

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func TestRemoveDirIfEmpty(t *testing.T) {
	type mocks struct {
		createDir  bool
		createFile bool
		readDirErr error
		removeErr  error
	}
	tests := []struct {
		name        string
		mocks       mocks
		shouldExist bool
	}{
		{
			name: "Should return error if ioutil.ReadDir raises it",
			mocks: mocks{
				createDir:  true,
				readDirErr: errors.New("some-read-dir-err"),
			},
			shouldExist: true,
		},
		{
			name: "Should return error if os.Remove raises it",
			mocks: mocks{
				createDir: true,
				removeErr: errors.New("some-remove-err"),
			},
			shouldExist: true,
		},
		{
			name: "Should remove dir if empty",
			mocks: mocks{
				createDir: true,
			},
			shouldExist: false,
		},
		{
			name: "Should not raise error if dir does not exist",
			mocks: mocks{
				createDir: false,
			},
			shouldExist: false,
		},
		{
			name: "Should not delete dir if not empty",
			mocks: mocks{
				createFile: true,
			},
			shouldExist: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := "not-existing-dir"
			if tt.mocks.createFile {
				path = filepath.Dir(testutils.TempFile(t, "some-file"))
			} else if tt.mocks.createDir {
				path = testutils.TempDir(t)
			}
			var wantErr error
			if tt.mocks.readDirErr != nil {
				wantErr = tt.mocks.readDirErr
				mockit.MockFunc(t, ioutil.ReadDir).With(path).Return(nil, wantErr)
			}
			if tt.mocks.removeErr != nil {
				wantErr = tt.mocks.removeErr
				mockit.MockFunc(t, os.Remove).With(path).Return(wantErr)
			}

			err := RemoveDirIfEmpty(path)

			assert.Equal(t, wantErr, err)
			if tt.shouldExist {
				testutils.PathExist(t, path)
			} else {
				testutils.PathDoesNotExist(t, path)
			}
		})
	}
}
