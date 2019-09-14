package prompt

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ParsePrompts parses the yaml file with the prompts definitions
func ParsePrompts(path string) ([]*Entry, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var prompts struct {
		Prompts []*Entry `yaml:"prompts,omitempty"`
	}
	yaml.Unmarshal(content, &prompts)

	// TODO: validate prompts

	return prompts.Prompts, nil
}
