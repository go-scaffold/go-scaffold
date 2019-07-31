package scaffold

import (
	"io/ioutil"
)

const templateSuffix = ".tpl"

func writeTextFile(content, outputFile string) error {
	return ioutil.WriteFile(outputFile, []byte(content), 0644)
}
