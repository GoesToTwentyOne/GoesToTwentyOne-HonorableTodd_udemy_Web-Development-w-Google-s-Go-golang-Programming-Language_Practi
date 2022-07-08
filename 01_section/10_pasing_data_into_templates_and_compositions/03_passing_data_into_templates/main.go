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

type Course struct {
	Code, Name, Unit string
}
type Semester struct {
	Term    string
	Courses []Course
}
type Year struct {
	Fall, Spring, Summer Semester
}

func main() {
	year := Year{
		Fall: Semester{
			Term: "Fall-1",
			Courses: []Course{
				{"101", "Banagla", "360"},
				{"102", "English", "360"},
				{"103", "DataBase Application", "360"},
				{"104", "DataBase Application", "360"},
				{"105", "Perifarals Device", "360"},
				{"106", "Operating Systems", "360"},
				{"106", "Data Structure and Algorithms", "360"},
				{"107", "PLC Control", "360"},
				{"108", "PLC Desing", "360"},
				{"109", "Microprocessor", "360"},
				{"110", "C#", "360"},
				{"111", "Java Aplication", "360"},
			},
		},
		Spring: Semester{
			Term: "Spring-1",
			Courses: []Course{
				{"101", "Banagla-2", "360"},
				{"102", "English-2", "360"},
				{"103", "DataBase Application-2", "360"},
				{"104", "DataBase Application-2", "360"},
				{"105", "Perifarals Device-2", "360"},
				{"106", "Operating Systems-2", "360"},
				{"106", "Data Structure and Algorithms-2", "360"},
				{"107", "PLC Control-2", "360"},
				{"108", "PLC Desing-2", "360"},
				{"109", "Microprocessor-2", "360"},
				{"110", "C#-2", "360"},
				{"111", "Java Aplication-2", "360"},
			},
		},
		Summer: Semester{
			Term: "Summer-1",
			Courses: []Course{
				{"101", "Banagla-3", "360"},
				{"102", "English-3", "360"},
				{"103", "DataBase Application-3", "360"},
				{"104", "DataBase Application-3", "360"},
				{"105", "Perifarals Device-3", "360"},
				{"106", "Operating Systems-3", "360"},
				{"106", "Data Structure and Algorithms-3", "360"},
				{"107", "PLC Control-3", "360"},
				{"108", "PLC Desing-3", "360"},
				{"109", "Microprocessor-3", "360"},
				{"110", "C#-3", "360"},
				{"111", "Java Aplication-3", "360"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", year)
	if err != nil {
		log.Fatal(err)
	}
}
