package strings

import (
	"strings"
)

func Replace(target string, replacement string, source string) string {
	return strings.ReplaceAll(source, target, replacement)
}
