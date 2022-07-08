package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tplr *template.Template

func init() {
	tplr = template.Must(template.New("").Funcs(ft).ParseFiles("tpl.gohtml"))

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

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var ft = template.FuncMap{
	"us":  strings.ToUpper,
	"ftl": firstThreeletter,
}

func firstThreeletter(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s

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
