package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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
	data := struct {
		Method     string
		Submission url.Values
	}{
		Method:     req.Method,
		Submission: req.Form,
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	var hd hotdog
	http.ListenAndServe(":5050", hd)

}
