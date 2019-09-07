package app

import (
	"log"
	"path/filepath"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/pasdam/go-scaffold/pkg/filter"
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

	prompts, err := prompt.ParsePrompts(filepath.Join(string(options.TemplatePath), ".go-scaffold", "prompts.yaml"))
	if err != nil {
		fatal("Unable to parse prompts.yaml file:", err)
		return
	}

	data := runPrompts(prompts)

	processOnlyTemplates := options.TemplatePath == options.OutputPath
	fileProcessor := scaffold.NewOutputFileProcessor(data, string(options.OutputPath), &scaffold.TemplateHelper{}, processOnlyTemplates)

	filter, _ := filter.NewPatternFilter(".go-scaffold(/.*)?")

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath))
	err = provider.ProvideFiles(filter, fileProcessor)
	if err != nil {
		fatal("Error while processing files:", err)
		return
	}
}
