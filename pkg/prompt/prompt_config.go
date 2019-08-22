package prompt

type PromptConfig struct {
	Name    string `yaml:"name,omitempty"`
	Type    string `yaml:"type,omitempty"`
	Default string `yaml:"default,omitempty"`
	Message string `yaml:"message,omitempty"`
}
