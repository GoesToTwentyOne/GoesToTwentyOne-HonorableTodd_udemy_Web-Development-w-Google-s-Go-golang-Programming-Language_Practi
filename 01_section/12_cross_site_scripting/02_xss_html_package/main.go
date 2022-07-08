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

type HTMLstructure struct {
	Title     string
	Heading   string
	Scripting string
}

func main() {
	ht := HTMLstructure{
		Title:     "Cross site scripting",
		Heading:   "This is first heading",
		Scripting: `<script>alert("Father of nation 100")</script>`,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ht)
	if err != nil {
		log.Fatal(err)
	}

}
