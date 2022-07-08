package main

import (
	"log"
	"os"
	"text/template"
)

type About struct {
	Name   string
	Age    int
	Expert string
}

var tpl *template.Template

func init() {

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	ab := About{Name: "Nihad", Age: 21, Expert: "golang"}
	err := tpl.Execute(os.Stdout, ab.Name)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", ab.Age)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", ab.Expert)
	if err != nil {
		log.Fatal(err)
	}

}
