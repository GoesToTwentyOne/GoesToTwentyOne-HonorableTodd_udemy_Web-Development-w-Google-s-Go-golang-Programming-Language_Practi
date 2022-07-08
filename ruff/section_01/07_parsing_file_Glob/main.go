package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("temp/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
