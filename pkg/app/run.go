package app

import (
	"log"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/prompt"
	"github.com/pasdam/go-scaffold/pkg/promptcli"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

var fatal = log.Fatal
var runPrompts func(prompts []*prompt.Entry) map[string]interface{} = promptcli.RunPrompts

// Run starts the app
func Run() {
	options, err := config.ParseCLIOptions()
	if err != nil {
		fatal("Command line options error:", err)
		return
	}

	prompts, err := prompt.ParsePrompts(options.PromptsConfigPath())
	if err != nil {
		fatal("Unable to parse prompts.yaml file:", err)
		return
	}

	data := runPrompts(prompts)

	processInPlace := options.TemplatePath == options.OutputPath
	templateHelper := &scaffold.TemplateHelper{}

	fileProcessor := scaffold.NewOutputFileProcessor(data, string(options.OutputPath), templateHelper, processInPlace)

	configFolderExcludeFilter, _ := filters.NewPatternFilter(false, "\\.go-scaffold(/.*)?")

	var fileToRemoveFilter filters.Filter
	if processInPlace && options.RemoveSource {
		fileToRemoveFilter = filters.Or(filters.NewPatternFilterFromInstance(configFolderExcludeFilter, true), templateHelper)
	}

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath), fileToRemoveFilter)
	err = provider.ProvideFiles(configFolderExcludeFilter, fileProcessor)
	if err != nil {
		fatal("Error while processing files. ", err)
		return
	}

	err = runInitScript(options.InitScriptPath(), string(options.OutputPath))
	if err != nil {
		fatal("Error while executing init script. ", err)
		return
	}
}
