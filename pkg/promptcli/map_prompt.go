package promptcli

import (
	"github.com/manifoldco/promptui"
	"github.com/pasdam/go-scaffold/pkg/prompt"
)

func mapPrompt(in *prompt.PromptConfig) *promptData {
	return &promptData{
		Prompt: &promptui.Prompt{
			Label:     in.Message,
			Default:   in.Default,
			IsConfirm: in.Type == "bool",
			Validate:  validateFuncForType(in.Type),
		},
		Name: in.Name,
	}
}
