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

type comparison struct {
	Num1 int
	Num2 int
}

func main() {
	c := []comparison{
		{2, 3},
		{3, 4},
		{4, 5},
		{20, 10},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", c)
	if err != nil {
		log.Fatal(err)
	}

}
