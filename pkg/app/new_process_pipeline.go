package app

import (
	"text/template"

	"github.com/pasdam/go-scaffold/pkg/core"
)

func newProcessPipeline(config interface{}, srcDir string, outDir string, errHandler func(v ...interface{}), funcMap template.FuncMap) core.Processor {
	outPipeline, err := newOutputPipeline(config, outDir, funcMap)
	if err != nil {
		errHandler("Error while creating the processing pipeline. ", err)
		return nil
	}
	return outPipeline
}
