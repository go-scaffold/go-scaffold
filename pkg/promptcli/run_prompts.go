package promptcli

import (
	"github.com/pasdam/go-scaffold/pkg/prompt"
)

var promptConfigToPromptUIMapper func(in *prompt.Entry) *promptData = mapPrompt

// RunPrompts executes the specified prompts
func RunPrompts(prompts []*prompt.Entry) map[string]interface{} {
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
