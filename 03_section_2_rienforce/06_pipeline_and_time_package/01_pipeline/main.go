package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fu).ParseFiles("tpl.gohtml"))
}

func double(i int) int {
	return i * 2
}
func sqrte(i int) int {
	return int(math.Sqrt(float64(i)))
}
func pow(i int) int {
	return int(math.Pow(float64(i), 2))
}

var fu = template.FuncMap{
	"d": double,
	"s": sqrte,
	"p": pow,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 21)
	if err != nil {
		log.Fatal(err)
	}

}
