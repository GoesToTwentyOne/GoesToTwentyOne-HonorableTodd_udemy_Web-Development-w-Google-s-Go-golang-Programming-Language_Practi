package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var h hotdog
	http.ListenAndServe(":5050", h)

}
