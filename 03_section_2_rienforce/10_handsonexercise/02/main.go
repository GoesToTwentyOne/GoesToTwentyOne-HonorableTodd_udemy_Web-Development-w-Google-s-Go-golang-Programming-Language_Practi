package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
	Name, Address, City, Zip string
}
type state struct {
	State  string
	Region string
	Hotels []hotel
}

func main() {
	s := []state{
		{
			State:  "California",
			Region: "Southern",
			Hotels: []hotel{
				{Name: "West End Hotel", Address: "Califonia-Fresno", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Mariana", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Fresno", Zip: "	90011"},
				{Name: "ARRIVE Palm Springs", Address: "Califonia-Los_Vagos", Zip: "	90011"},
			},
		},
		{
			State:  "California",
			Region: "Central",
			Hotels: []hotel{
				{Name: "Comfort Inn Carme", Address: "Califonia-Central", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Mariana", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Fresno", Zip: "	90011"},
				{Name: "ARRIVE Palm Springs", Address: "Califonia-Los_Vagos", Zip: "	90011"},
			},
		},
		{
			State:  "California",
			Region: "Northern",
			Hotels: []hotel{
				{Name: "Comfort Inn Carme", Address: "Califonia-Central", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Mariana", Zip: "	90011"},
				{Name: "Bay Club Hotel and Marina", Address: "Califonia-Fresno", Zip: "	90011"},
				{Name: "ARRIVE Palm Springs", Address: "Califonia-Los_Vagos", Zip: "	90011"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)
	if err != nil {
		log.Fatal(err)
	}
}
