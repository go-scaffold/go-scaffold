package prompt_test

import (
	"testing"

	"github.com/pasdam/go-project-template/pkg/prompt"
	"github.com/stretchr/testify/assert"
)

func Test_ParsePrompts_Success_ShouldParsePromptsIfFileExists(t *testing.T) {
	prompts, err := prompt.ParsePrompts("testdata/prompts.yaml")

	assert.Nil(t, err)
	assert.Equal(t, 3, len(prompts))
	assert.Equal(t, &prompt.PromptConfig{
		Name:    "prompt1",
		Type:    "string",
		Default: "prompt1-default-val",
		Message: "Enter prompt1 value",
	}, prompts[0])
	assert.Equal(t, &prompt.PromptConfig{
		Name:    "prompt2",
		Type:    "bool",
		Default: "true",
		Message: "Enter prompt2 value",
	}, prompts[1])
	assert.Equal(t, &prompt.PromptConfig{
		Name:    "prompt3",
		Type:    "int",
		Default: "prompt3-default-val",
		Message: "Enter prompt3 value",
	}, prompts[2])
}

func Test_ParsePrompts_Fail_ShouldReturnErrorIfFileDoesNotExists(t *testing.T) {
	prompts, err := prompt.ParsePrompts("not-existing-yaml")

	assert.NotNil(t, err)
	assert.Nil(t, prompts)
}
