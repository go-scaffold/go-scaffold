package prompt

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParsePrompts(path string) ([]*PromptConfig, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var prompts struct {
		Prompts []*PromptConfig `yaml:"prompts,omitempty"`
	}
	yaml.Unmarshal(content, &prompts)

	// TODO: validate prompts

	return prompts.Prompts, nil
}
