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

type CompanyEmployee struct {
	Name   string
	Expert string
	Admin  bool
}

func main() {
	cE := []CompanyEmployee{
		{Name: "Nihad Hossain", Expert: "golang", Admin: true},
		{Name: "Aneta  Lonora", Expert: "golang", Admin: false},
		{Name: "Cyneya Hossain", Expert: "golang", Admin: true},
		{Name: "Kenlala ", Expert: "golang", Admin: false},
		{Name: "Anny Hossain", Expert: "golang", Admin: true},
		{Name: "Holritoin Trump", Expert: "golang", Admin: false},
		{Name: "Angela Duck Worth", Expert: "golang", Admin: true},
		{Name: "Nitlami Reba", Expert: "golang", Admin: false},
		{Name: "Herry Melton", Expert: "golang", Admin: true},
		{Name: "Arel Xelmol", Expert: "golang", Admin: true},
		{Name: " ", Expert: "golang", Admin: true},
		{Name: "Antyu White", Expert: "golang", Admin: true},
		{Name: "Htyu Cripton", Expert: "golang", Admin: true},
		{Name: "Junior  Hytion", Expert: "golang", Admin: true},
		{Name: "Gelton Mane", Expert: "golang", Admin: true},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", cE)
	if err != nil {
		log.Fatal(err)
	}

}
