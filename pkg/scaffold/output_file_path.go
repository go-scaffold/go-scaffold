package scaffold

import (
	"strings"
)

func OutputFilePath(filePath string) string {
	if IsTemplate(filePath) {
		return strings.TrimSuffix(filePath, templateSuffix)
	}
	return filePath
}
