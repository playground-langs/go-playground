package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	text := "Template start\n Action: {{.}}\n Template end\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)
	err = tmpl.Execute(os.Stdout, 42)
	check(err)
	err = tmpl.Execute(os.Stdout, true)
	check(err)

	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true!{{end}} finish\n", false)
	t := "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(t, []float64{1.25, 0.99, 27})
	executeTemplate(t, nil)

	s := "Name:{{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	executeTemplate(s, Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true})
	executeTemplate(s, Subscriber{Name: "Joy Carr", Rate: 5.99, Active: false})
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}
