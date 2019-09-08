package scaffold

import (
	"bytes"
	"log"
	"text/template"
)

func ApplyTemplate(templateContent string, config interface{}) (string, error) {
	template, err := template.New("").Parse(templateContent)
	if err != nil {
		log.Println("Error while parsing template")
		return "", err
	}

	var result bytes.Buffer
	err = template.Execute(&result, config)
	if err != nil {
		log.Println("Error generating content from template")
		return "", err
	}

	return result.String(), nil
}
