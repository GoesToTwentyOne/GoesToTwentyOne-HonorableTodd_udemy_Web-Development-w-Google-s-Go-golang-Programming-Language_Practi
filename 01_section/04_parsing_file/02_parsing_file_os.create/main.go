package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
