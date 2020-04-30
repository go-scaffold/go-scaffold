package app

import (
	"errors"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/go-scaffold/pkg/testutils"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_newOutputPipeline_ShouldReturnErrorIfOneOccursWhenCreatingTheFilter(t *testing.T) {
	type mocks struct {
		filterInclusive bool
		filterPattern   string
	}
	type args struct {
		inPlace bool
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
	}{
		{
			name: "Should return error if the process is in place",
			mocks: mocks{
				filterInclusive: false,
				filterPattern:   "\\.(go-scaffold|git)(" + filepath.FromSlash("/") + ".*)?",
			},
			args: args{
				inPlace: false,
			},
		},
		{
			name: "Should return error if the process is not in place",
			mocks: mocks{
				filterInclusive: true,
				filterPattern:   "\\.*\\.tpl",
			},
			args: args{
				inPlace: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := errors.New("some-error")
			var pattern string
			captor := func(v interface{}) bool {
				pattern = v.([]string)[0]
				return true
			}
			mockit.MockFunc(t, filters.NewPatternFilter).With(tt.mocks.filterInclusive, captor).Return(nil, wantErr)

			got, err := newOutputPipeline(tt.args.inPlace, nil, "", nil)

			assert.Equal(t, tt.mocks.filterPattern, pattern)
			assert.Nil(t, got)
			assert.Equal(t, wantErr, err)
		})
	}
}

func Test_newOutputPipeline_ShouldProcessFileAsExpected(t *testing.T) {
	data := make(map[string]interface{})
	data["text"] = "test!"
	outDir := testutils.TempDir(t)
	type args struct {
		inPlace bool
		path    string
		content string
	}
	type expect struct {
		shouldExist bool
		path        string
	}
	tests := []struct {
		name   string
		args   args
		expect expect
	}{
		{
			name: "Should ignore go-scaffold files if the process is in place",
			args: args{
				inPlace: true,
				path:    filepath.Join(".go-scaffold", "prompts.yaml"),
				content: "",
			},
			expect: expect{
				shouldExist: false,
				path:        filepath.Join(".go-scaffold", "prompts.yaml"),
			},
		},
		{
			name: "Should ignore regular files if the process is in place",
			args: args{
				inPlace: true,
				path:    "some-regular-file.txt",
				content: "",
			},
			expect: expect{
				shouldExist: false,
				path:        "some-regular-file.txt",
			},
		},
		{
			name: "Should process template files if the process is in place",
			args: args{
				inPlace: true,
				path:    "some-template-file.txt.tpl",
				content: "some-template-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-template-file.txt",
			},
		},

		{
			name: "Should ignore go-scaffold files if the process is not in place",
			args: args{
				inPlace: false,
				path:    filepath.Join(".go-scaffold", "prompts.yaml"),
				content: "",
			},
			expect: expect{
				shouldExist: false,
				path:        filepath.Join(".go-scaffold", "prompts.yaml"),
			},
		},
		{
			name: "Should process regular files if the process is not in place",
			args: args{
				inPlace: false,
				path:    "some-regular-file.txt",
				content: "some-regular-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-regular-file.txt",
			},
		},
		{
			name: "Should process template files if the process is not in place",
			args: args{
				inPlace: false,
				path:    "some-template-file.txt.tpl",
				content: "some-template-file-content",
			},
			expect: expect{
				shouldExist: true,
				path:        "some-template-file.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newOutputPipeline(tt.args.inPlace, data, outDir, &scaffold.TemplateHelper{})
			assert.NotNil(t, got)
			assert.Nil(t, err)

			reader := strings.NewReader(tt.args.content)

			err = got.ProcessFile(tt.args.path, reader)
			assert.Nil(t, err)

			outPath := filepath.Join(outDir, tt.expect.path)
			if tt.expect.shouldExist {
				testutils.FileExistsWithContent(t, outPath, tt.args.content)
			} else {
				testutils.PathDoesNotExist(t, outPath)
			}
		})
	}
}
