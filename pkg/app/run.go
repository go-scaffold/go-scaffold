package app

import (
	"log"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/pasdam/go-scaffold/pkg/prompt"
	"github.com/pasdam/go-scaffold/pkg/promptcli"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
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

	filter, _ := scaffold.NewPatternFilter(".go-scaffold(/.*)?")

	provider, err := scaffold.NewFileSystemProvider(string(options.TemplatePath), filter)
	if err != nil {
		fatal("Error while creating the file provider:", err)
		return
	}

	prompts, err := prompt.ParsePrompts(filepath.Join(string(options.TemplatePath), ".go-scaffold", "prompts.yaml"))
	if err != nil {
		fatal("Unable to parse prompts.yaml file:", err)
		return
	}

	data := runPrompts(prompts)


	processOnlyTemplates := options.TemplatePath == options.OutputPath
	err = scaffold.ProcessFiles(provider, data, string(options.OutputPath), processOnlyTemplates)
	if err != nil {
		fatal("Error while processing files:", err)
		return
	}
}
