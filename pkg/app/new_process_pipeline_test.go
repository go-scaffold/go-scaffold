package app

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/otiai10/copy"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_newProcessPipeline_ShouldNotIncludeCleanupPipelineIfProcessIsNotInPlace(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!!"
	outDir := testutils.TempDir(t)
	errHandler := func(v ...interface{}) {
		t.Fail() // errors should not occur
	}
	srcDir := filepath.Join("testdata", "valid_template")
	got := newProcessPipeline(false, data, srcDir, outDir, &scaffold.TemplateHelper{}, errHandler)

	path := "file.txt.tpl"
	got.ProcessFile(path, strings.NewReader("This is a {{ .text }}\n"))

	testutils.FileExistsWithContent(t, filepath.Join(srcDir, path), "This is a {{ .text }}\n")
	testutils.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!!\n")
}

func Test_newProcessPipeline_ShouldIncludeCleanupPipelineIfProcessIsInPlace(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!!"
	outDir := testutils.TempDir(t)
	errHandler := func(v ...interface{}) {
		t.Fail() // errors should not occur
	}
	srcDir := filepath.Join("testdata", "valid_template")
	got := newProcessPipeline(true, data, outDir, outDir, &scaffold.TemplateHelper{}, errHandler)
	copy.Copy(srcDir, outDir)

	path := "file.txt.tpl"
	got.ProcessFile(path, strings.NewReader("This is a {{ .text }}\n"))

	testutils.PathDoesNotExist(t, filepath.Join(outDir, path))
	testutils.FileExistsWithContent(t, filepath.Join(outDir, "file.txt"), "This is a test!!\n")
}

func Test_newProcessPipeline_ShouldReturnErrorIfOneOccurWhileCreatingThePipelines(t *testing.T) {
	type mocks struct {
		outPipelineErr   error
		cleanPipelineErr error
	}
	type args struct {
		inPlace bool
		config  interface{}
		srcDir  string
		outDir  string
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
				inPlace: false,
				config:  "some-out-pipeline-error-data",
				srcDir:  "some-out-pipeline-error-source-dir",
				outDir:  "some-out-pipeline-error-output-dir",
			},
		},
		{
			name: "Should return error if one occurs while creating cleanup pipeline",
			mocks: mocks{
				cleanPipelineErr: errors.New("some-cleanup-pipeline-error"),
			},
			args: args{
				inPlace: true,
				config:  "some-cleanup-pipeline-error-data",
				srcDir:  "some-cleanup-pipeline-error-source-dir",
				outDir:  "some-cleanup-pipeline-error-output-dir",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helper := &scaffold.TemplateHelper{}
			wantErr := tt.mocks.outPipelineErr
			errorOccurred := false
			errHandler := func(v ...interface{}) {
				errorOccurred = true

				assert.Equal(t, "Error while creating the processing pipeline. ", v[0])
				assert.Equal(t, wantErr, v[1])
			}

			mockit.MockFunc(t, newOutputPipeline).With(tt.args.inPlace, tt.args.config, tt.args.outDir, helper).Return(nil, tt.mocks.outPipelineErr)
			mockit.MockFunc(t, newCleanupPipeline).With(tt.args.srcDir).Return(nil, tt.mocks.cleanPipelineErr)
			if tt.mocks.cleanPipelineErr != nil {
				wantErr = tt.mocks.cleanPipelineErr
			}

			got := newProcessPipeline(tt.args.inPlace, tt.args.config, tt.args.srcDir, tt.args.outDir, helper, errHandler)

			assert.True(t, errorOccurred)
			assert.Nil(t, got)
		})
	}
}
