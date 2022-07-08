package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fu).ParseFiles("tpl.gohtml"))
}

func format(t time.Time) string {
	return t.Format("Mon Jan _2 15:04:05 2006")
}

var fu = template.FuncMap{
	"f": format,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}

}
