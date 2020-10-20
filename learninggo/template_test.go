package learninggo

import (
	"os"
	"testing"
	"text/template"
)

type Part struct {
	Name  string
	Count int
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

func TestTemplates(t *testing.T) {

	templatesMoreCoplex()
}

func templatesMoreCoplex() {
	templateText := "Name: {{.Name}}\n {{if .Active}} Rate: {{.Rate}} {{end}}\n"
	exectueTemplate(templateText, Subscriber{Name: "Homer Simpson", Active: false, Rate: 23.2})
	exectueTemplate(templateText, Subscriber{Name: "Ned  Flanders", Active: true, Rate: 2.6})
}

func loopingStructs() {
	templateText := "Name: {{.Name}} -- Count: {{.Count}}\n"
	exectueTemplate(templateText, Part{Name: "some Part", Count: 23})
	exectueTemplate(templateText, Part{Name: "Other Part", Count: 26})
}

func exectueTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func firstTemplate() {
	exectueTemplate("Dot is {{.}}\n", "ABC")
	exectueTemplate("Dot is {{.}}\n", 123.5)
}

func ifTemplate() {

	exectueTemplate("Dot is {{if .}}true {{end}}finsh\n", true)
	exectueTemplate("Dot is {{if .}}true {{end}}finsh\n", false)
}

func loopingTemplate() {
	data := []string{"do", "re", "mi"}
	exectueTemplate("Befor loop {{.}}\n{{range .}} In loop: {{.}}\n{{end}}After loop: {{.}}\nfinsh\n", data)

	prices := []float64{1.23, 4.78, 3.52}
	exectueTemplate("Prices\n{{range .}}{{.}}\n{{end}}\n", prices)

}
