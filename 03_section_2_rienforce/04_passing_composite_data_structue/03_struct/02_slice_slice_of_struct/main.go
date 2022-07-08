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

func main() {
	ye := []year{
		{
			AcademicYear: "2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"66553", "The Introduction of Golang", "26"},
					{"66773", "Rich Dad Poor Dad", "10"},
					{"67676", "As A Man Thinketh", "7"},
					{"66645", "Essential Programming Golang", "12"},
					{"66684", "Ultimate Go", "45"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					{"66553", "The Introduction of Golang", "26"},
					{"66773", "Rich Dad Poor Dad", "10"},
					{"67676", "As A Man Thinketh", "7"},
					{"66645", "Essential Programming Golang", "12"},
					{"66684", "Ultimate Go", "45"},
				},
			},
		},
		{
			AcademicYear: "2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					{"66553", "The Introduction of Golang", "26"},
					{"66773", "Rich Dad Poor Dad", "10"},
					{"67676", "As A Man Thinketh", "7"},
					{"66645", "Essential Programming Golang", "12"},
					{"66684", "Ultimate Go", "45"},
				},
			},
			Summer: semester{
				Term: "Summer",
				Courses: []course{
					{"66553", "The Introduction of Golang", "26"},
					{"66773", "Rich Dad Poor Dad", "10"},
					{"67676", "As A Man Thinketh", "7"},
					{"66645", "Essential Programming Golang", "12"},
					{"66684", "Ultimate Go", "45"},
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, ye)
	if err != nil {
		log.Fatal(err)
	}

}
