package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("first.gmao")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}
	tpl, err = tpl.ParseFiles("second.gmao", "third.gmao")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "third.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "second.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

}
