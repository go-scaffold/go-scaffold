package helpers

import (
	"bytes"
	"log"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gertd/go-pluralize"
	"github.com/go-scaffold/go-sdk/v2/pkg/templates"
)

func TemplateFunctions(funcMaps ...template.FuncMap) template.FuncMap {
	sprigFuncMap := sprig.TxtFuncMap()

	customFuncs := make(template.FuncMap)
	pluralizeClient := pluralize.NewClient()
	customFuncs["pluralize"] = pluralizeClient.Plural
	customFuncs["singularize"] = pluralizeClient.Singular
	customFuncs["isPlural"] = pluralizeClient.IsPlural
	customFuncs["isSingular"] = pluralizeClient.IsSingular
	customFuncs["debug"] = func(v ...any) string {
		log.Printf("[TEMPLATE DEBUG] %v", v)
		return "" // Return empty string to not affect template output
	}

	maps := make([]template.FuncMap, 0, len(funcMaps)+2)
	maps = append(maps, sprigFuncMap)
	maps = append(maps, customFuncs)
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
