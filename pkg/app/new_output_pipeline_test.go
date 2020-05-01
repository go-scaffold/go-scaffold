package app

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func Test_ignorePattern(t *testing.T) {
	regEx, err := regexp.Compile(ignorePattern)
	assert.Nil(t, err)

	tests := []struct {
		name        string
		shouldMatch bool
	}{
		{name: ".git", shouldMatch: true},
		{name: ".git" + string(os.PathSeparator) + "some-file", shouldMatch: true},
		{name: ".gitignore", shouldMatch: false},
		{name: ".go-scaffold", shouldMatch: true},
		{name: ".go-scaffold" + string(os.PathSeparator) + "some-file", shouldMatch: true},
		{name: ".go-scaffold-file", shouldMatch: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.shouldMatch, regEx.MatchString(tt.name))
		})
	}
}

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
				filterPattern:   "\\.(go-scaffold|git)(" + string(os.PathSeparator) + ".*)?$",
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
	outDir := filestest.TempDir(t)
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
			name: "Should process files with .go-scaffold prefix",
			args: args{
				inPlace: false,
				path:    ".go-scaffold-file",
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        ".go-scaffold-file",
			},
		},
		{
			name: "Should ignore .git folder",
			args: args{
				inPlace: false,
				path:    filepath.Join(".git", "some-git-file"),
				content: "",
			},
			expect: expect{
				shouldExist: false,
				path:        filepath.Join(".git", "some-git-file"),
			},
		},
		{
			name: "Should process files with .git prefix",
			args: args{
				inPlace: false,
				path:    ".gitignore",
				content: "",
			},
			expect: expect{
				shouldExist: true,
				path:        ".gitignore",
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
		{
			name: "Should not process go-scaffold config files",
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
				filestest.FileExistsWithContent(t, outPath, tt.args.content)
			} else {
				filestest.PathDoesNotExist(t, outPath)
			}
		})
	}
}
