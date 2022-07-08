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

type Person struct {
	Name string
	Age  int
}
type SecretAgent struct {
	Person
	LTK bool
}

func main() {
	// pe := SecretAgent{
	// 	Person: Person{
	// 		Name: "Nihad Hossain",
	// 		Age:  21,
	// 	},
	// 	LTK: false,
	// }
	pe2 := SecretAgent{
		Person: Person{
			Name: "Xelton",
			Age:  21,
		},
		LTK: true,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pe2)
	if err != nil {
		log.Fatal(err)
	}

	// err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", pe)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
