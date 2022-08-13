package app

import "text/template"

func mergeFuncMaps(maps ...template.FuncMap) template.FuncMap {
	res := make(template.FuncMap)
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
