package config_test

import (
	"os"
	"testing"

	"github.com/pasdam/go-scaffold/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_ParseCLIOption_fail_shouldReturnErrorIfAnInvalidParameterIsSpecifiedSpecified(t *testing.T) {
	templateDir := "some-template-dir"
	mockArguments(false, templateDir, "", true)
	os.Args = append(os.Args, "--invalid-param")

	options, err := config.ParseCLIOption()

	assert.NotNil(t, err)
	assert.Equal(t, "unknown flag `invalid-param'", err.Error())
	assert.Nil(t, options)
}

func Test_ParseCLIOption_success_shouldUseDefaultOutputPath(t *testing.T) {
	templateDir := "some-template-dir"
	mockArguments(false, templateDir, "", true)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, "./", string(options.OutputPath))
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.True(t, options.RemoveConfig)
}

func Test_ParseCLIOption_success_shouldUseDefaultRemoveConfig(t *testing.T) {
	outDir := "some-output-dir"
	templateDir := "some-template-dir"
	mockArguments(false, templateDir, outDir, false)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.False(t, options.RemoveConfig)
}

func Test_ParseCLIOption_success_shouldUseDefaultTemplatePath(t *testing.T) {
	outDir := "some-output-dir"
	mockArguments(false, "", outDir, true)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.Equal(t, "./", string(options.TemplatePath))
	assert.True(t, options.RemoveConfig)
}

func Test_ParseCLIOption_success_shouldParseOptionsWithLongFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	mockArguments(true, templateDir, outDir, true)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.True(t, options.RemoveConfig)
}

func Test_ParseCLIOption_success_shouldParseOptionsWithShortFlags(t *testing.T) {
	templateDir := "some-template-dir"
	outDir := "some-output-dir"
	mockArguments(false, templateDir, outDir, true)

	options, err := config.ParseCLIOption()

	assert.Nil(t, err)
	assert.NotNil(t, options)
	assert.Equal(t, templateDir, string(options.TemplatePath))
	assert.Equal(t, outDir, string(options.OutputPath))
	assert.True(t, options.RemoveConfig)
}

func mockArguments(useLongFlags bool, templateDir string, outDir string, withRemoveConfig bool) {
	os.Args = make([]string, 7)
	os.Args[0] = ""
	if templateDir != "" {
		if useLongFlags {
			os.Args[1] = "--template"
		} else {
			os.Args[1] = "-t"
		}
		os.Args[2] = templateDir
	}
	if outDir != "" {
		if useLongFlags {
			os.Args[3] = "--output"
		} else {
			os.Args[3] = "-o"
		}
		os.Args[4] = outDir
	}
	if withRemoveConfig {
		if useLongFlags {
			os.Args[5] = "--remove-config"
		} else {
			os.Args[5] = "-r"
		}
		os.Args[6] = outDir
	}
}
