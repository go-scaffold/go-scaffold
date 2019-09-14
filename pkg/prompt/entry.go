package prompt

// Entry contains the configuration of a prompt instance
type Entry struct {

	// Name is the name of the variable in the template replaced with the input value of this prompt
	Name string `yaml:"name,omitempty"`

	// Type indicates the variable type
	Type string `yaml:"type,omitempty"`

	// Default indicates the value to use if the prompt is not inserted
	Default string `yaml:"default,omitempty"`

	// Message is the message to show while prompting for the value
	Message string `yaml:"message,omitempty"`
}
