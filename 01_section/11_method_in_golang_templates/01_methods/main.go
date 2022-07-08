package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) AgePrint() int {
	return p.Age

}
func (p Person) DoubleAgePrint() int {
	return p.Age * 2

}
func (p Person) TkeAgePrint(x int) int {
	return x * 2

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	pre := Person{
		Name: "Nihad Hossain",
		Age:  21,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pre)
	if err != nil {
		log.Fatal(err)
	}

}
