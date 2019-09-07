package scaffold

import (
	"strings"
)

const templateSuffix = ".tpl"

type TemplateHelper struct{}

func (self *TemplateHelper) Accept(filePath string) bool {
	return strings.HasSuffix(filePath, templateSuffix)
}

func (self *TemplateHelper) OutputFilePath(filePath string) string {
	return strings.TrimSuffix(filePath, templateSuffix)
}
