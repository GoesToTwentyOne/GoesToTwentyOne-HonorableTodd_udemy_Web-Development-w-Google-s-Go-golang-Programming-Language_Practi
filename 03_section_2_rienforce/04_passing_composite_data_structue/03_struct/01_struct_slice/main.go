package main

import (
	"html/template"
	"log"
	"os"
)

type course struct {
	Code, Name, Units string
}
type semester struct {
	Term    string
	Courses []course
}
type year struct {
	AcademicYear         string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ye := year{
		AcademicYear: "2021",
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"101", "Bangla", "14"},
				{"101", "English", "14"},
				{"103", "Math", "14"},
				{"103", "Math", "14"},
				{"66634", "Data Structure and Algorithom", "12"},
				{"103", "C#", "12"},
				{"103", "Golang", "15"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, ye)
	if err != nil {
		log.Fatal(err)
	}

}
