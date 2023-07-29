package helpers

import (
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/go-scaffold/go-scaffold/pkg/helpers/collections"
	"github.com/go-scaffold/go-scaffold/pkg/helpers/strings"
	"github.com/iancoleman/strcase"
)

func TemplateFunctions(funcMaps ...template.FuncMap) template.FuncMap {
	sprigFuncMap := sprig.TxtFuncMap()
	sprigFuncMap["camelcase"] = strcase.ToCamel
	sprigFuncMap["replace"] = strings.Replace
	sprigFuncMap["sequence"] = collections.Sequence
	maps := make([]template.FuncMap, 0, len(funcMaps)+1)
	maps = append(maps, sprigFuncMap)
	maps = append(maps, funcMaps...)
	return mergeFuncMaps(maps)
}
