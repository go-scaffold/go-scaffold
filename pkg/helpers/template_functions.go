package helpers

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/go-scaffold/go-scaffold/pkg/helpers/collections"
	"github.com/go-scaffold/go-scaffold/pkg/helpers/strings"
	"github.com/go-scaffold/go-sdk/v2/pkg/templates"
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

func TemplateAwareFunctions() templates.TemplateAwareFuncMap {
	maps := templates.TemplateAwareFuncMap{
		"include": func(tpl *template.Template) any {
			return func(name string, data any) (string, error) {
				var result bytes.Buffer
				err := tpl.ExecuteTemplate(&result, name, data)
				if err != nil {
					return "", err
				}
				return result.String(), nil
			}
		},
	}
	return maps
}
