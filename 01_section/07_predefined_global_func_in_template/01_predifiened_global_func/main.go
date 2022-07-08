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
	xs := []string{"Annnt", "Annya", "Rito London", "Yello Gino", "Alex Goot", "Gelion Funfu", "Nihad"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xs)
	if err != nil {
		log.Fatal(err)
	}

}
