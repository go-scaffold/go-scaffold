package app

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_newProcessPipeline_ShouldProcessFilesCorrectly(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!!"
	outDir := filestest.TempDir(t)
	errHandler := func(v ...interface{}) {
		t.Fail() // errors should not occur
	}
	got := newProcessPipeline(data, outDir, errHandler, nil)

	path := "file.txt"
	got.ProcessFile(path, strings.NewReader("This is a {{ .text }}\n"))

	filestest.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!!\n")
}

func Test_newProcessPipeline_ShouldReturnErrorIfOneOccurWhileCreatingThePipelines(t *testing.T) {
	type mocks struct {
		outPipelineErr   error
		cleanPipelineErr error
	}
	type args struct {
		config interface{}
		outDir string
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
	}{
		{
			name: "Should return error if one occurs while creating out pipeline",
			mocks: mocks{
				outPipelineErr: errors.New("some-out-pipeline-error"),
			},
			args: args{
				config: "some-out-pipeline-error-data",
				outDir: "some-out-pipeline-error-output-dir",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := tt.mocks.outPipelineErr
			errorOccurred := false
			errHandler := func(v ...interface{}) {
				errorOccurred = true

				assert.Equal(t, "Error while creating the processing pipeline. ", v[0])
				assert.Equal(t, wantErr, v[1])
			}

			mockit.MockFunc(t, newOutputPipeline).With(tt.args.config, tt.args.outDir, nil).Return(nil, tt.mocks.outPipelineErr)
			if tt.mocks.cleanPipelineErr != nil {
				wantErr = tt.mocks.cleanPipelineErr
			}

			got := newProcessPipeline(tt.args.config, tt.args.outDir, errHandler, nil)

			assert.True(t, errorOccurred)
			assert.Nil(t, got)
		})
	}
}
