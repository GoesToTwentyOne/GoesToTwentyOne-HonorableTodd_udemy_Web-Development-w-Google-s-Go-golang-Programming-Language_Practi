package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	comperison := struct {
		ComOneValue int
		ComTwoValue int
	}{
		ComOneValue: 45,
		ComTwoValue: 48,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", comperison)
	if err != nil {
		log.Fatal(err)
	}

}
