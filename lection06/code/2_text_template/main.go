package main

import (
	"os"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Actions []string
}

func main() {
	d := easyTemplate()
	err := d.Execute(os.Stdout, Person{Name: "Arnold"})
	if err != nil {
		panic(err)
	}

	d = advancedTemplate()
	err = d.Execute(os.Stdout, []Person{{Name: "Arnold"}, {Name: "Justin", Age: 35, Actions: []string{"eat", "sleep", "learn", "repeat"}}})
	if err != nil {
		panic(err)
	}
}

func easyTemplate() *template.Template {
	const easyT = "Hello, {{.Name}}!"

	return template.Must(template.New("easy").Parse(easyT))
}

func advancedTemplate() *template.Template {
	const advancedT = `
{{ range . }}
Name: {{ .Name }}
{{ if .Age -}} Age: {{ .Age }} {{ else -}} Age is undefined {{- end }}
{{ range .Actions -}} Action is {{ . }}
{{ else -}} No Actions detected
{{ end }} {{ end }}
`
	return template.Must(template.New("advanced").Parse(advancedT))
}
