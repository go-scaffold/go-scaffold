package config_test

import (
	"os"
	"testing"

	"github.com/pasdam/go-project-template/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_ParseCLIOption_fail_shouldReturnErrorIfNoArgumentIsSpecified(t *testing.T) {
	options, err := config.ParseCLIOption()

	assert.NotNil(t, err)
	assert.Equal(t, "the required flag `-o, --output' was not specified", err.Error())
	assert.Nil(t, options)
}

func Test_ParseCLIOption_fail_shouldReturnErrorIfOutputDirIsNotSpecified(t *testing.T) {
	templateDir := "some-template-dir"
	mockArguments(false, &templateDir, nil)

	options, err := config.ParseCLIOption()

	assert.NotNil(t, err)
	assert.Equal(t, "the required flag `-o, --output' was not specified", err.Error())
	assert.Nil(t, options)
}

func Test_ParseCLIOption_success_shouldUseDefaultTemplatePath(t *testing.T) {
	outDir := "some-output-dir"
	mockArguments(false, nil, &outDir)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, "./", string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
}

func Test_ParseCLIOption_success_shouldParseTemplatePathAndOutPathWithLongFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	mockArguments(true, &templateDir, &outDir)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
}

func Test_ParseCLIOption_success_shouldParseTemplatePathAndOutPathWithShortFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	mockArguments(false, &templateDir, &outDir)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
}

func mockArguments(useLongFlags bool, templateDir *string, outDir *string) {
	os.Args = make([]string, 5)
	os.Args[0] = ""
	if templateDir != nil {
		if useLongFlags {
			os.Args[1] = "--template"
		} else {
			os.Args[1] = "-t"
		}
		os.Args[2] = *templateDir
	}
	if outDir != nil {
		if useLongFlags {
			os.Args[3] = "--output"
		} else {
			os.Args[3] = "-o"
		}
		os.Args[4] = *outDir
	}
}
