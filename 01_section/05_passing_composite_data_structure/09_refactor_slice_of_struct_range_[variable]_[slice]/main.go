package main

import (
	"log"
	"os"
	"text/template"
)

var tplr *template.Template

func init() {
	tplr = template.Must(template.ParseFiles("tpl.gohtml"))

}

type Cars struct {
	Sedan      int
	COUPE      int
	SPORTS_CAR int
	Customers  string
}
type Bycycle struct {
	Road        int
	Mountain    int
	Quantity    int
	CustomersIn string
}

func main() {
	ca := []Cars{
		{25551, 2772, 15, "American"},
		{25551, 2772, 15, "American"},
		{251, 252, 15, "American"},
		{251, 22, 15, "American"},
		{2551, 252, 15, "American"},
		{251, 22, 15, "American"},
		{201, 252, 15, "American"},
		{201, 22, 15, "American"},
		{210, 22, 15, "American"},
		{26, 262, 15, "American"},
		{26, 262, 15, "American"},
		{22, 262, 15, "American"},
		{24, 272, 15, "American"},
		{26, 200, 15, "American"},
		{25, 2882, 15, "American"},
	}
	bi := []Bycycle{
		{25551, 2772, 15, "India"},
		{25551, 2772, 15, "India"},
		{251, 252, 15, "India"},
		{251, 22, 15, "India"},
		{2551, 252, 15, "India"},
		{251, 22, 15, "India"},
		{201, 252, 15, "India"},
		{201, 22, 15, "India"},
		{210, 22, 15, "India"},
		{26, 262, 15, "India"},
		{26, 262, 15, "India"},
		{22, 262, 15, "India"},
		{24, 272, 15, "India"},
		{26, 200, 15, "India"},
		{25, 2882, 15, "India"},
	}
	data := struct {
		Ca []Cars
		Bi []Bycycle
	}{
		Ca: ca,
		Bi: bi,
	}
	err := tplr.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
