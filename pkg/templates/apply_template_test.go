package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyTemplate_Fail_ShouldReturnErrorIfItFailsToExecuteTheTemplate(t *testing.T) {
	result, err := applyTemplate("This is a {{ .NotExistingProperty }}", "invalid_config")

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_applyTemplate_Fail_ShouldReturnErrorIfTemplateIsInvalid(t *testing.T) {
	data := struct{ CustomProperty string }{CustomProperty: "*test*"}
	result, err := applyTemplate("This is a {{ .CustomProperty } with invalid template", data)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_applyTemplate_Success_ShouldCorrectlyGenerateOutputContentFromTemplate(t *testing.T) {
	data := struct{ CustomProperty string }{CustomProperty: "*test*"}
	result, err := applyTemplate("This is a {{ .CustomProperty }}", data)

	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*", result)
}
