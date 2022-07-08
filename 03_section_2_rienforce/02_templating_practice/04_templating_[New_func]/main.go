package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "Nihad Hossain",
		Age:  21,
	}

	//parsing a file and temporary name
	tpl, err := template.New("").Parse("My name is {{.Name}}.I am {{.Age}}")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, person)
	if err != nil {
		log.Fatal(err)
	}

}
