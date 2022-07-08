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

type Person struct {
	Name string
	Age  int
}

func main() {
	per := Person{
		Name: "Nihad",
		Age:  21,
	}

	// anoPer := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "Seb Linux",
	// 	Age:  22,
	// }
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", per)
	if err != nil {
		log.Fatal(err)
	}

}
