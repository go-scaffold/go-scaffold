package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-scaffold/go-scaffold/pkg/config"
)

func TestOptions_ManifestPath(t *testing.T) {
	type fields struct {
		ManifestName     string
		OutputPath       string
		TemplateRootPath string
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Should return expected value with default manifest",
			fields: fields{
				TemplateRootPath: "manifest-test",
			},
			want: filepath.Join("manifest-test", "Manifest.yaml"),
		},
		{
			name: "Should return expected value with custom manifest",
			fields: fields{
				ManifestName:     "Chart.yaml",
				TemplateRootPath: "chart-test",
			},
			want: filepath.Join("chart-test", "Chart.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				ManifestName:     tt.fields.ManifestName,
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.ManifestPath(); got != tt.want {
				t.Errorf("Options.ManifestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_TemplateDirPath(t *testing.T) {
	type fields struct {
		OutputPath       string
		TemplateRootPath string
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplateRootPath: "template-test"},
			want:   filepath.Join("template-test", "templates"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.TemplateDirPath(); got != tt.want {
				t.Errorf("Options.TemplateDirPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptions_ValuesPath(t *testing.T) {
	type fields struct {
		OutputPath       string
		TemplateRootPath string
		Values           []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Should return expected value",
			fields: fields{TemplateRootPath: "values-test"},
			want:   filepath.Join("values-test", "values.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &config.Options{
				OutputPath:       tt.fields.OutputPath,
				TemplateRootPath: tt.fields.TemplateRootPath,
				Values:           tt.fields.Values,
			}
			if got := o.ValuesPath(); got != tt.want {
				t.Errorf("Options.ValuesPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockArguments(useLongFlags bool, templateDir string, outDir string) []string {
	oldArgs := os.Args

	os.Args = make([]string, 7)
	os.Args[0] = ""
	if templateDir != "" {
		if useLongFlags {
			os.Args[1] = "--template"
		} else {
			os.Args[1] = "-t"
		}
		os.Args[2] = templateDir
	}
	if outDir != "" {
		if useLongFlags {
			os.Args[3] = "--output"
		} else {
			os.Args[3] = "-o"
		}
		os.Args[4] = outDir
	}

	return oldArgs
}
