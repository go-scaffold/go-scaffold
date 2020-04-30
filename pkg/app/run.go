package app

import (
	"log"

	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

var errHandler = log.Fatal

// Run starts the app
func Run() {
	options := readOptions(errHandler)
	data := parseAndRunPrompts(options.PromptsConfigPath(), errHandler)

	processInPlace := options.TemplatePath == options.OutputPath
	templateHelper := &scaffold.TemplateHelper{}

	fileProcessor := newProcessPipeline(
		processInPlace,
		data,
		string(options.TemplatePath),
		string(options.OutputPath),
		templateHelper,
		errHandler,
	)

	provider := scaffold.NewFileSystemProvider(string(options.TemplatePath))
	err := provider.ProvideFiles(nil, fileProcessor)
	if err != nil {
		errHandler("Error while processing files. ", err)
		return
	}

	runInitScript(options.InitScriptPath(), string(options.OutputPath), errHandler)
}
