package passwd

import (
	"io"
	"text/template"
)

var (
	tmpl template.Template
)

type File struct {
	Users []User
}

type User struct {
	Name  string
	UID   string
	GID   string
	Home  string
	Shell string
}

func Write(f File, w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "passwd", f)
}

func init() {
	t, err := template.New("passwd").Parse(`
{{define "passwd" -}}
{{range .Users -}}
{{.Name}}:x:{{.UID}}:{{.UID}}::{{.Home}}:{{.Shell}}
{{end}}
{{- end}}
	`)

	if err != nil {
		panic(err.Error())
	}

	tmpl = *t
}
