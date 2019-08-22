package promptcli

import (
	"github.com/pasdam/go-project-template/pkg/prompt"
)

var promptConfigToPromptUIMapper func(in *prompt.PromptConfig) *promptData = mapPrompt

func RunPrompts(prompts []*prompt.PromptConfig) map[string]interface{} {
	result := make(map[string]interface{})
	for _, config := range prompts {
		prompt := promptConfigToPromptUIMapper(config)
		val, _ := prompt.Prompt.Run()
		// TODO: handle error
		// TODO: convert into proper type
		result[prompt.Name] = val
	}
	return result
}
