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

	processInPlace := options.TemplatePath == options.OutputPath
	templateHelper := &scaffold.TemplateHelper{}

	fileProcessor := scaffold.NewOutputFileProcessor(data, string(options.OutputPath), templateHelper, processInPlace)

	configFolderExcludeFilter, _ := filter.NewPatternFilter(false, "\\.go-scaffold(/.*)?")

	var fileToRemoveFilter filter.Filter
	if processInPlace && options.RemoveSource {
		fileToRemoveFilter = filter.NewMultiFilter(configFolderExcludeFilter.NewInstance(true), templateHelper)
	}

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath), fileToRemoveFilter)
	err = provider.ProvideFiles(configFolderExcludeFilter, fileProcessor)
	if err != nil {
		fatal("Error while processing files:", err)
		return
	}
}
