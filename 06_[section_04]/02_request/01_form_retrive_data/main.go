package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	var d hotdog
	http.ListenAndServe(":5050", d)

}
