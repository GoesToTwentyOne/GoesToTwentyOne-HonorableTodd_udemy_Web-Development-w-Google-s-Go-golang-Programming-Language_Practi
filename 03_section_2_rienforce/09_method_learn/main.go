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

type double struct {
	Odd  int
	Even int
	Age  int
}

func (d double) Odddouble() int {
	return d.Odd * 2

}
func (d double) Evendouble(x int) int {
	return d.Even * 2

}
func (d double) Agedouble(x int) int {
	return d.Age * 2

}

func main() {
	do := double{
		Odd:  5,
		Even: 10,
		Age:  21,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", do)
	if err != nil {
		log.Fatal(err)
	}

}
