# Named template for project header
{{- define "partials.header" -}}
# {{ .Values.projectName }}
{{ .Values.description }}

**Author:** {{ .Values.author }}
**Version:** {{ .Values.version }}

---

{{- end -}}