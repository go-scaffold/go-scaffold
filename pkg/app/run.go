package app

import (
	"log"

	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

var errHandler = log.Fatal

// Run starts the app
func Run() {
	options := readOptions(errHandler)
	data := parseAndRunPrompts(options.PromptsConfigPath(), errHandler)

	processInPlace := options.TemplatePath == options.OutputPath
	templateHelper := &scaffold.TemplateHelper{}

	fileProcessor := scaffold.NewOutputFileProcessor(data, string(options.OutputPath), templateHelper, processInPlace)

	configFolderExcludeFilter, _ := filters.NewPatternFilter(false, "\\.go-scaffold(/.*)?")

	var fileToRemoveFilter filters.Filter
	if processInPlace && options.RemoveSource {
		fileToRemoveFilter = filters.Or(filters.NewPatternFilterFromInstance(configFolderExcludeFilter, true), templateHelper)
	}

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath), fileToRemoveFilter)
	err := provider.ProvideFiles(configFolderExcludeFilter, fileProcessor)
	if err != nil {
		errHandler("Error while processing files. ", err)
		return
	}

	runInitScript(options.InitScriptPath(), string(options.OutputPath), errHandler)
}
