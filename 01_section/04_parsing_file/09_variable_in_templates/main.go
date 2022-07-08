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
	err := tpl.Execute(os.Stdout, "take time and try to understand master wisdom")
	if err != nil {
		log.Fatal(err)
	}

}
