package template_package

import (
	"os"
	"strings"
	"text/template"
)

type User struct{
	Name		string
	Description	string
}

type Users struct {
	Users	[]User
}

func TemplateTesting() {
	// String
	tmpString := "Hello {{.}}\n"
	tmpVar := "Adebayo"
	tmpAge := 21
	t := template.Must(template.New("hello").Parse(tmpString))
	t.Execute(os.Stdout, tmpVar)
	t.Execute(os.Stdout, tmpAge)

	tmpString = "{{ range . }}Hello {{ . }}\n{{end}}"
	tmpArray := []string{"Clin", "Ade", "Joeboy"}
	t = template.Must(template.New("array").Parse(tmpString))
	t.Execute(os.Stdout, tmpArray)

	tmpStruct := []User{
		{Name: "Clinton", Description: "Good Boy"},
		{Name: "Adebayo", Description: "Cool Boy"},
		{Name: "Ade", Description: "Finest"},
	}

	tmpString = `{{range . }} 
	Name: {{ .Name }} Description: {{ .Description }} {{end}}`
	t = template.Must(template.New("struct").Parse(tmpString))
	t.Execute(os.Stdout, tmpStruct)

	users := Users{
		Users: []User{
			{Name: "Okeke", Description: "Okeke is a girl"},
			{Name: "Joeboy", Description: "A singer"},
		},
	}
	
	

	tmpString = `{{with .Users }}
	users:
	{{range .}}
	Name: {{toUpper .Name }}, Description: {{ .Description}}

	{{end}}
	{{end}}`

	funcMap := map[string]interface{} {
		"toUpper": strings.ToUpper,
	}
	t = template.Must(template.New("advance").Funcs(funcMap).Parse(tmpString))
	t.Execute(os.Stdout, users)

	tmpString = `{{with .Users }}
	users:
	{{range $index, $element := .}}
	{{increment $index}}. Name: {{ $element.Name }}, Description: {{ $element.Description}}
	{{end}}
	{{end}}`
	
	funcMap = template.FuncMap{
		"increment": func (i int) int  {
			return i + 1
		},
	}
	t = template.Must(template.New("advance").Funcs(funcMap).Parse(tmpString))
	t.Execute(os.Stdout, users)
	
}