package promptcli

import (
	"testing"

	"github.com/pasdam/go-scaffold/pkg/prompt"

	"github.com/stretchr/testify/assert"
)

type mockPrompt struct {
	val string
}

func (m *mockPrompt) Run() (string, error) {
	return m.val, nil
}

func mockMapper(in *prompt.Entry) *promptData {
	return &promptData{
		Name:   in.Name,
		Prompt: &mockPrompt{val: in.Name + "-val"},
	}
}

func TestRunPrompts(t *testing.T) {
	promptConfigToPromptUIMapper = mockMapper

	prompts := make([]*prompt.Entry, 3)
	prompts[0] = &prompt.Entry{
		Default: "dp0",
		Type:    "string",
		Message: "mp0",
		Name:    "p0",
	}
	prompts[1] = &prompt.Entry{
		Default: "dp1",
		Type:    "int",
		Message: "mp1",
		Name:    "p1",
	}
	prompts[2] = &prompt.Entry{
		Default: "f",
		Type:    "bool",
		Message: "mp2",
		Name:    "p2",
	}

	data := RunPrompts(prompts)

	assert.Equal(t, "p0-val", data["p0"])
	assert.Equal(t, "p1-val", data["p1"])
	assert.Equal(t, "p2-val", data["p2"])
}
