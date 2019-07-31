package scaffold

import (
	"strings"
)

func IsTemplate(filePath string) bool {
	return strings.HasSuffix(filePath, templateSuffix)
}
