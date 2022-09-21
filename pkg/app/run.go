package app

import (
	"log"
	"text/template"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/pasdam/go-scaffold/pkg/providers"
	"github.com/pasdam/go-template-map-loader/pkg/tm"
)

var errHandler = log.Fatal

// Run starts the app
func Run(options *config.Options, funcMap template.FuncMap) {
	if options.TemplateRootPath == options.OutputPath {
		log.Fatal("Can't generate file in the input folder, please specify an output directory")
		return
	}

	manifestData, err := tm.LoadYamlFile(options.ManifestPath())
	if err != nil {
		log.Fatal("An error occurred while reading the manifest file: ", err.Error())
		return
	}

	defaultValuesData, err := tm.LoadYamlFile(options.ValuesPath())
	if err != nil {
		log.Fatal("An error occurred while reading the values file: ", err.Error())
		return
	}

	valuesData := make([]map[string]interface{}, 0, len(options.Values)+1)
	valuesData = append(valuesData, defaultValuesData)

	for _, path := range options.Values {
		data, err := tm.LoadYamlFile(path)
		if err != nil {
			log.Fatal("An error occurred while reading the values file: ", err.Error())
			return
		}
		valuesData = append(valuesData, data)
	}

	data := tm.MergeMaps(
		tm.WithPrefix("Manifest", manifestData),
		tm.WithPrefix("Values", tm.MergeMaps(valuesData...)),
	)

	fileProcessor := newProcessPipeline(
		data,
		string(options.OutputPath),
		errHandler,
		funcMap,
	)

	provider := providers.NewFileSystemProvider(string(options.TemplateDirPath()))
	err = provider.ProvideFiles(nil, fileProcessor)
	if err != nil {
		errHandler("Error while processing files. ", err)
		return
	}
}
