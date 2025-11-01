package config_test

import (
	"path/filepath"
	"testing"

	"github.com/go-scaffold/go-scaffold/pkg/config"
)

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
