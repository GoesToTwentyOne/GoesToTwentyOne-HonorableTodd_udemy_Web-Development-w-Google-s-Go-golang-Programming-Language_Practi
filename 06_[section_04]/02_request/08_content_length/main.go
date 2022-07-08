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
		Method        string
		URL           *url.URL
		Submission    url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		Method:        req.Method,
		URL:           req.URL,
		Submission:    req.Form,
		Header:        req.Header,
		Host:          req.Host,
		ContentLength: req.ContentLength,
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
