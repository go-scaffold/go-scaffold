package scaffold_test

import (
	"testing"

	"github.com/pasdam/go-project-template/pkg/scaffold"
	"github.com/stretchr/testify/assert"
)

func Test_ApplyTemplate_Fail_ShouldReturnErrorIfItFailsToExecuteTheTemplate(t *testing.T) {
	result, err := scaffold.ApplyTemplate("This is a {{ .NotExistingProperty }}", "invalid_config")

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_ApplyTemplate_Fail_ShouldReturnErrorIfTemplateIsInvalid(t *testing.T) {
	data := struct{ CustomProperty string }{CustomProperty: "*test*"}
	result, err := scaffold.ApplyTemplate("This is a {{ .CustomProperty } with invalid template", data)

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func Test_ApplyTemplate_Success_ShouldCorrectlyGenerateOutputContentFromTemplate(t *testing.T) {
	data := struct{ CustomProperty string }{CustomProperty: "*test*"}
	result, err := scaffold.ApplyTemplate("This is a {{ .CustomProperty }}", data)

	assert.Nil(t, err)
	assert.Equal(t, "This is a *test*", result)
}
