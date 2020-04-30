package app

import (
	"github.com/pasdam/go-scaffold/pkg/processors"
	"github.com/pasdam/go-scaffold/pkg/scaffold"
)

func newProcessPipeline(inPlace bool, config interface{}, srcDir string, outDir string, templateHelper *scaffold.TemplateHelper, errHandler func(v ...interface{})) processors.Processor {
	outPipeline, err := newOutputPipeline(inPlace, config, outDir, templateHelper)
	if err != nil {
		errHandler("Error while creating the processing pipeline. ", err)
		return nil
	}

	if inPlace {
		cleanupPipeline, err := newCleanupPipeline(srcDir)
		if err != nil {
			errHandler("Error while creating the processing pipeline. ", err)
			return nil
		}

		return processors.NewPipeline(outPipeline, cleanupPipeline)
	}

	return outPipeline
}
