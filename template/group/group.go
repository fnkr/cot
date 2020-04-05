package group

import (
	"io"
	"strings"
	"text/template"
)

var (
	tmpl template.Template
)

type File struct {
	Groups []Group
}

type Group struct {
	Name    string
	GID     string
	Members []string
}

func Write(f File, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "group", f)
}

func init() {
	t := template.New("group")
	t.Funcs(template.FuncMap{"JoinStrings": strings.Join})
	if _, err := t.Parse(`
{{define "group" -}}
{{range .Groups -}}
{{.Name}}:x:{{.GID}}:{{ JoinStrings .Members "," }}
{{end}}
{{- end}}
	`); err != nil {
		panic(err.Error())
	}

	tmpl = *t
}
