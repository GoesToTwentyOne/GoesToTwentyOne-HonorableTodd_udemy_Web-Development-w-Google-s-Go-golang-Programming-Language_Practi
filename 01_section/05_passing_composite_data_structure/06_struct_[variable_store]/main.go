package main

import (
	"log"
	"os"
	"text/template"
)

type BevrageStore struct {
	Lager     int
	Ale       int
	Stout     int
	Customers string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	bev := BevrageStore{
		Lager:     240,
		Ale:       21,
		Stout:     69,
		Customers: "first class citizen",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", bev)
	if err != nil {
		log.Fatal(err)
	}

}
