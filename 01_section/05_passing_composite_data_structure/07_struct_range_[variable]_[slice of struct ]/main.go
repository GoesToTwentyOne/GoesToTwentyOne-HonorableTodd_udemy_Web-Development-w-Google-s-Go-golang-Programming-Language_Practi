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
	bev := []BevrageStore{
		{230, 45, 23, "Cheif Secretary"},
		{230, 45, 23, "Cheif UFO"},
		{230, 45, 23, "Cheif Vice Pre"},
		{230, 45, 23, "Cheif Presedent"},
		{230, 45, 23, "Cheif Secretary"},
		{230, 45, 23, "Cheif Secretary General"},
		{230, 45, 23, "Cheif Secretary Head"},
		{230, 45, 23, "Deputy Cheif Secretary "},
		{230, 45, 23, "Deputy Cheif Secretary General"},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", bev)
	if err != nil {
		log.Fatal(err)
	}

}
