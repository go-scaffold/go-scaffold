package config_test

import (
	"testing"

	"github.com/pasdam/go-project-template/pkg/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewTemplateConfig_fail_shouldReturnErrorIfFailsToOpenFile(t *testing.T) {
	templateConfig, err := config.NewTemplateConfig("test/not-existing-file.yaml")

	assert.NotNil(t, err)
	assert.Nil(t, templateConfig)
}

func Test_NewTemplateConfig_success_shouldCreateConfigIfFileIsValid(t *testing.T) {
	templateConfig, err := config.NewTemplateConfig("test/test_config.yaml")

	assert.Nil(t, err)
	assert.NotNil(t, templateConfig)

	config := templateConfig.Config()
	group1Config := config["group1"].(map[string]interface{})
	assert.NotNil(t, config)
	assert.Equal(t, "val1", config["key1"])
	assert.Equal(t, "val2", config["key2"])
	assert.Equal(t, "*test*", config["Text"])
	assert.Equal(t, "valA", group1Config["keyA"])
	assert.Equal(t, "valB", group1Config["keyB"])
}
