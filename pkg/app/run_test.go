package app

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/go-scaffold/go-scaffold/pkg/config"
	"github.com/pasdam/go-files-test/pkg/filestest"
	"github.com/pasdam/go-utils/pkg/assertutils"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	type args struct {
		options  *config.Options
		funcMaps []template.FuncMap
	}
	type wantFile struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    []*wantFile
		wantErr error
	}{
		{
			name: "Should process valid template",
			args: args{
				options: &config.Options{
					TemplateRootPath: filepath.Join("testdata", "valid_template"),
				},
			},
			want: []*wantFile{
				{
					path:    "file.txt",
					content: "This is a test!\n",
				},
				{
					path:    "normal_file.txt",
					content: "normal-file-content\n",
				},
				{
					path:    "service-a",
					content: "config: some-config-service-a\n",
				},
				{
					path:    "service-b",
					content: "config: some-config-service-b\n",
				},
			},
		},
		{
			name: "Should return error if output dir is same as input",
			args: args{
				options: &config.Options{
					OutputPath:       "testdata",
					TemplateRootPath: "testdata",
				},
			},
			wantErr: errors.New("can't generate file in the input folder, please specify an output directory"),
		},
		{
			name: "Should return error if manifest is not found",
			args: args{
				options: &config.Options{
					TemplateRootPath: filepath.Join("testdata", "invalid_templates"),
				},
			},
			wantErr: errors.New("an error occurred while reading the manifest file: open testdata/invalid_templates/Manifest.yaml: no such file or directory"),
		},
		{
			name: "Should return error if default values file does not exist",
			args: args{
				options: &config.Options{
					TemplateRootPath: filepath.Join("testdata", "invalid_templates", "no_values"),
				},
			},
			wantErr: errors.New("error while loading data: open testdata/invalid_templates/no_values/values.yaml: no such file or directory"),
		},
		{
			name: "Should return error if custom values file does not exist",
			args: args{
				options: &config.Options{
					TemplateRootPath: filepath.Join("testdata", "valid_template"),
					Values: []string{
						"some-not-existing-file",
					},
				},
			},
			wantErr: errors.New("error while loading data: open some-not-existing-file: no such file or directory"),
		},
		{
			name: "Should return error if processing the templates raises one",
			args: args{
				options: &config.Options{
					TemplateRootPath: filepath.Join("testdata", "invalid_templates", "invalid_syntax"),
				},
			},
			wantErr: errors.New("error while running the pipeline: template: :2: unclosed action started at :1"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertOutDir := false
			if tt.args.options.OutputPath == "" {
				outDir := filestest.TempDir(t)
				tt.args.options.OutputPath = outDir
				assertOutDir = true
			}

			err := Run(tt.args.options, tt.args.funcMaps...)

			assertutils.AssertEqualErrors(t, tt.wantErr, err)
			if assertOutDir {
				entries, err := os.ReadDir(tt.args.options.OutputPath)
				assert.NoError(t, err)
				assert.Equal(t, len(tt.want), len(entries))
			}
			for _, file := range tt.want {
				filestest.FileExistsWithContent(t, filepath.Join(tt.args.options.OutputPath, file.path), file.content)
			}
		})
	}
}

func TestRunWithFileProvider_ShouldReturnErrorIfTemplateProviderIsNil(t *testing.T) {
	options := &config.Options{
		TemplateRootPath: filepath.Join("testdata", "invalid_templates", "empty_values"),
	}

	err := RunWithFileProvider(options, nil)

	assert.Error(t, err)
	assert.Equal(t, "error while building the processing pipeline: no template processor specified for the pipeline", err.Error())
}
