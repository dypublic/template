package main

import (
	"os"

	"template_light/template"
)

type Friend struct {
	Fname string
}
type HomeTown struct {
	City string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
	HomeTown HomeTown
}

func main() {
	f1 := Friend{Fname: "xiaofang"}
	f2 := Friend{Fname: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(`
hello {{.UserName}} at {{.HomeTown.City}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Fname}}
{{- end }}
{{ end }}`))
	//{{define "x"}}x{{end}}
	//{{template "y"}}`
	p := Person{UserName: "longshuai",
		Emails:   []string{"a1@qq.com", "a2@gmail.com"},
		Friends:  []*Friend{&f1, &f2},
		HomeTown: HomeTown{"sh"},
	}

	t.Execute(os.Stdout, p)
}
