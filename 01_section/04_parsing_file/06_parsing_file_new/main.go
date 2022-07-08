package main

import (
	"html/template"
	"log"
	"os"
)

type AboutME struct {
	Name        string
	ExpertField string
}

func main() {
	about := AboutME{
		Name:        "Nihad Hossain",
		ExpertField: "Golang Web dev",
	}
	tpl, err := template.New("abouts").Parse("My name is \"{{.Name}}\"  I'm expart the golang tool \"{{.ExpertField}}\" ")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(os.Stdout, about)
	if err != nil {
		log.Fatal(err)
	}

}
