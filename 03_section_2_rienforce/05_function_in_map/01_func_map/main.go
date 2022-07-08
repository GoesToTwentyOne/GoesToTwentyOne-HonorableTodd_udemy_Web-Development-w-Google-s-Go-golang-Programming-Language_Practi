package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(functem).ParseFiles("tpl.gohtml"))
	// tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

var functem = template.FuncMap{

	"uppercase": strings.ToUpper,
}

type person struct {
	Name string
	Age  int
}

func main() {
	pe := []person{
		{Name: "Nihad", Age: 21},
		{Name: "Wandun", Age: 19},
		{Name: "Hendul", Age: 18},
		{Name: "Yenrewa", Age: 12},
		{Name: "Seb", Age: 21},
		{Name: "Helageui", Age: 34},
		{Name: "Momdek", Age: 32},
		{Name: "Dedulia", Age: 21},
		{Name: "Sumaliya", Age: 14},
		{Name: "Neha", Age: 07},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pe)
	if err != nil {
		log.Fatal(err)
	}

}
