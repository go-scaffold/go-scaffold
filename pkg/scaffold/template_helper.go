package scaffold

import (
	"strings"
)

const templateSuffix = ".tpl"

// TemplateHelper is a type with helper methods used to process template files
type TemplateHelper struct{}

// Accept returns true if the specified file is a template, false otherwise
func (p *TemplateHelper) Accept(filePath string) bool {
	return strings.HasSuffix(filePath, templateSuffix)
}

// OutputFilePath returns the name of the output file, given the name of the input one
func (p *TemplateHelper) OutputFilePath(filePath string) string {
	return strings.TrimSuffix(filePath, templateSuffix)
}
