package app

import (
	"github.com/pasdam/go-scaffold/pkg/prompt"
	"github.com/pasdam/go-scaffold/pkg/promptcli"
)

func parseAndRunPrompts(promptsConfigPath string, errHandler func(v ...interface{})) map[string]interface{} {
	prompts, err := prompt.ParsePrompts(promptsConfigPath)
	if err != nil {
		errHandler("Unable to parse prompts.yaml file. ", err)
		return nil
	}

	return promptcli.RunPrompts(prompts)
}
