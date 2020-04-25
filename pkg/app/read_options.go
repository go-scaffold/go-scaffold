package app

import "github.com/pasdam/go-scaffold/pkg/config"

func readOptions(errHandler func(v ...interface{})) *config.Options {
	options, err := config.ParseCLIOptions()
	if err != nil {
		errHandler("Command line options error. ", err)
		return nil
	}
	return options
}
