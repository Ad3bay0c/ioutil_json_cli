package template_package

import (
	"fmt"
	"os"
	"text/template"
)

type Todo struct {
	Name		string
	Description	string
}
func TemplateT() {
	var todo = Todo {
		Name: "Working on a project",
		Description: "Command Line Interface Project using Cobra",
	}

	t, _ := template.New("todo").Parse("The Title: {{ .Name }} and Description: {{ .Description }}\n")
	t.Execute(os.Stdout, todo)

	t1 := template.Must(template.New("todo").Parse("We Move so let's go {{ .Description }}\n"))
	t1.Execute(os.Stdout, todo)

	b := []byte("<h4>First Heading</h4><br><p>This is a paragraph</p>")
	template.HTMLEscape(os.Stdout, b)
	fmt.Println(template.HTMLEscapeString("Okay, let's see</>"))
	
	fmt.Println(template.IsTrue("Ok"))

	fmt.Println(template.JSEscapeString("<script>console.log('123')</script>"))

	story := ` Hi, {{ .Name }}; how are you doing, been a long time.
	You were invited for our reunion sometimes ago and 
	{{if .Attended}} it was my pleasure to have seen your handsome face once again.
	{{- else}} it was a shame for not showing up for the meeting. {{- end}}
	{{with .Gift -}} Thank you and we really appreciate with a {{.}}.{{end}}
	Best Wishes`
	var s = struct {
		Name,Gift	string
		Attended	bool
	}{
		Name: "Segun",
		Gift: "Car",
		Attended: true,
	}
	var s1 = struct {
		Name,Gift	string
		Attended	bool
	}{
		Name: "Segun",
		Gift: "Car",
		Attended: false,
	}

	t2 := template.Must(template.New("story").Parse(story))
	t2.Execute(os.Stdout, s)
	t2.Execute(os.Stdout, s1)

	fmt.Println(t2.DefinedTemplates())
}