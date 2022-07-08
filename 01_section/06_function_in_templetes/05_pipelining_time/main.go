package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(tF).ParseFiles("tpl.gohtml"))
}
func timeF(t time.Time) string {
	return t.Format("01-12-1999")
}

var tF = template.FuncMap{
	"tFormat": timeF,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}

}
