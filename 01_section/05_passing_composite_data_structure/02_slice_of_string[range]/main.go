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
	xs := []string{"Nihad", "Neha", "Anyel", "Seb"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Fatal(err)
	}

}
