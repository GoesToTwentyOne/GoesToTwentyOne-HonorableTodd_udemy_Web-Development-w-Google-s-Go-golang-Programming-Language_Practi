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
	Ltk  bool
}

func main() {
	pe := []person{
		{Name: "Nihad Hossain", Age: 21, Ltk: true},
		{Name: "Anyel", Age: 19, Ltk: false},
		{Name: "Hreyuy", Age: 18, Ltk: false},
		{Name: "", Age: 12, Ltk: true},
		{Name: "Sadia", Age: 25, Ltk: false},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pe)
	if err != nil {
		log.Fatal(err)
	}

}
