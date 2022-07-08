package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("temp/*"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}

}
