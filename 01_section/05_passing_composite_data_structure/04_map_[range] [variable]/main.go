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
	m := map[string]string{
		"Name":      "Nihad",
		"Address":   "Hazipara",
		"Education": "Golang Focus Time",
		"Eye":       "Two eye clear",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", m)
	if err != nil {
		log.Fatal(err)
	}

}
