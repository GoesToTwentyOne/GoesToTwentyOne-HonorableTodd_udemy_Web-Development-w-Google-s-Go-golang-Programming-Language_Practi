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

func main() {
	err := tpl.Execute(os.Stdout, []int{1, 2, 3, 4, 5, 6, 7, 89})
	if err != nil {
		log.Fatal(err)
	}

}
