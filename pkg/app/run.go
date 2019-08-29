package app

import (
	"log"
	"path/filepath"

	"github.com/pasdam/go-project-template/pkg/prompt"
	"github.com/pasdam/go-project-template/pkg/promptcli"

	"github.com/pasdam/go-project-template/pkg/config"
	"github.com/pasdam/go-project-template/pkg/scaffold"
)

var fatal = log.Fatal
var runPrompts func(prompts []*prompt.PromptConfig) map[string]interface{} = promptcli.RunPrompts

// Run starts the app
func Run() {
	options, err := config.ParseCLIOption()
	if err != nil {
		fatal("Command line options error:", err)
		return
	}

	// TODO: check if template path contains a template

	prompts, err := prompt.ParsePrompts(filepath.Join(string(options.TemplatePath), "prompts.yaml"))
	if err != nil {
		fatal("Unable to parse prompts.yaml file:", err)
		return
	}

	data := runPrompts(prompts)

	provider, err := scaffold.NewFileSystemProvider(filepath.Join(string(options.TemplatePath), "src"))
	if err != nil {
		fatal("Error while creating the file provider:", err)
		return
	}

	err = scaffold.ProcessFiles(provider, data, string(options.OutputPath), false)
	if err != nil {
		fatal("Error while processing files:", err)
		return
	}
}
