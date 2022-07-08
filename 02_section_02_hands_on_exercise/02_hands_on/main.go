package main

import (
	"log"
	"os"
	"text/template"
)

type aboutCalifonia struct {
	Name, Address, City, Region string
	Zip                         int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	ab := []aboutCalifonia{
		{
			Name:    "Nihad Hossain",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Dallas",
			Region:  "Southern",
			Zip:     75247,
		},
		{
			Name:    "Dear Anyel",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Dallas",
			Region:  "Northern",
			Zip:     75247,
		},
		{
			Name:    "Jertyullika",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Houseton",
			Region:  "NorHousetonthern",
			Zip:     75247,
		},
		{
			Name:    "Bobe Hanry",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Houseton",
			Region:  "Southern",
			Zip:     75247,
		},
		{
			Name:    "Johb Melton  ",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Dallas",
			Region:  "Southern",
			Zip:     75247,
		},
		{
			Name:    "Asyle Waikins",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Houseton",
			Region:  "Central",
			Zip:     75247,
		},
		{
			Name:    "Rtyudury",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Austin",
			Region:  "Southern",
			Zip:     75247,
		},
		{
			Name:    "Ertyruliya",
			Address: "44 EAST AVEAUSTIN, TX 78701",
			City:    "Austin",
			Region:  "Central",
			Zip:     75247,
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ab)
	if err != nil {
		log.Fatal(err)
	}

}
