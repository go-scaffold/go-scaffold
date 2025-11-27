package config

import (
	"path/filepath"
)

// Options contains the app run configuration
type Options struct {
	OutputPath       string
	TemplateRootPath string
	Values           []string
	SkipUnchanged    bool
}

// TemplateDirPath returns the path of the template dir
func (o *Options) TemplateDirPath() string {
	return filepath.Join(string(o.TemplateRootPath), "templates")
}
