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

type person struct {
	Name string
	Age  int
}

func main() {
	pe := []person{

		{"Nihad", 21},
		{"Ahona", 34},
		{"Fibni", 23},
		{"Grto", 35},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pe)

	if err != nil {
		log.Fatal(err)
	}

}
