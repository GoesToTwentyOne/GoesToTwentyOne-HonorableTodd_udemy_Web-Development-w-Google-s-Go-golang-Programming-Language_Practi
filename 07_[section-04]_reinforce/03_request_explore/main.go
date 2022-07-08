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
		Header        http.Header
		ContentLength int64
		Host          string
		Form          url.Values
		RemoteAddr    string
	}{
		Method:        req.Method,
		URL:           req.URL,
		Header:        req.Header,
		ContentLength: req.ContentLength,
		Host:          req.Host,
		Form:          req.Form,
		RemoteAddr:    req.RemoteAddr,
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var h hotdog
	err := http.ListenAndServe(":5050", h)
	if err != nil {
		log.Fatal(err)
	}

}
