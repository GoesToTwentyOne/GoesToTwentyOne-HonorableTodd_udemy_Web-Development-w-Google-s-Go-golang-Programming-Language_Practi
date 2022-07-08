package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl := template.Must(template.ParseGlob("template/*"))
	err := tpl.ExecuteTemplate(os.Stdout, "first.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

}
