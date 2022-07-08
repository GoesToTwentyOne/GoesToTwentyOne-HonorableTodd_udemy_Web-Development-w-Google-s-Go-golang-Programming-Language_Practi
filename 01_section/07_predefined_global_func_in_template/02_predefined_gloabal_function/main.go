package main

import (
	"html/template"
	"log"
	"os"
)

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	xs := []string{"Annnt", "Annya", "Rito London", "Yello Gino", "Alex Goot", "Gelion Funfu", "Nihad"}
	na := struct {
		Name []string
		Age  int
	}{
		Name: xs,
		Age:  21,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", na)
	if err != nil {
		log.Fatal(err)
	}

}
