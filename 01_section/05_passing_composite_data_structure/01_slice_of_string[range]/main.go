package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	name := []string{"Nihad", "Neha", "Anyel", "Seba", "Todd"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", name)
	if err != nil {
		log.Fatal(err)
	}

}
