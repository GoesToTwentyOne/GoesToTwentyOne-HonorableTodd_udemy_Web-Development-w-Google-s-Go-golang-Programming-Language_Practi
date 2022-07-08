package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
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
	data := struct {
		Method     string
		URL        *url.URL
		Submission url.Values
	}{
		Method:     req.Method,
		URL:        req.URL,
		Submission: req.Form,
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	var h hotdog
	http.ListenAndServe(":5050", h)

}
