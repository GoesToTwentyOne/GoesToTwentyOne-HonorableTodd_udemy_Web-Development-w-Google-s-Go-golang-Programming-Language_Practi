package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Nihad")
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}

}
