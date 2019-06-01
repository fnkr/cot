package group

import (
	"io"
	"text/template"
)

var (
	tmpl template.Template
)

type File struct {
	Groups []Group
}

type Group struct {
	Name string
	GID  string
}

func Write(f File, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "group", f)
}

func init() {
	t, err := template.New("group").Parse(`
{{define "group" -}}
{{range .Groups -}}
{{.Name}}:x:{{.GID}}:
{{end}}
{{- end}}
	`)

	if err != nil {
		panic(err.Error())
	}

	tmpl = *t
}
