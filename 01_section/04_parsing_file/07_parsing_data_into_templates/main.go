package main

import (
	"html/template"
	"log"
	"os"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	err := t.Execute(os.Stdout, "Nihad Hossain")
	if err != nil {
		log.Fatal(err)
	}

}
