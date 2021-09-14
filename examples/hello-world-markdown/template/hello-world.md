# Project {{ .Values.project }}

{{ .Values.intro }}

## Links

{{ range .Values.links -}}
- [{{ .name }}]({{ .url }})
{{ end -}}
