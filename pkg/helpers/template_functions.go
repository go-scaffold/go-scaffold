package helpers

import (
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/iancoleman/strcase"
	"github.com/pasdam/go-scaffold/pkg/helpers/collections"
	"github.com/pasdam/go-scaffold/pkg/helpers/strings"
)

func TemplateFunctions(funcMap template.FuncMap) template.FuncMap {
	sprigFuncMap := sprig.TxtFuncMap()
	sprigFuncMap["camelcase"] = strcase.ToCamel
	sprigFuncMap["replace"] = strings.Replace
	sprigFuncMap["sequence"] = collections.Sequence
	return mergeFuncMaps(sprigFuncMap, funcMap)
}
