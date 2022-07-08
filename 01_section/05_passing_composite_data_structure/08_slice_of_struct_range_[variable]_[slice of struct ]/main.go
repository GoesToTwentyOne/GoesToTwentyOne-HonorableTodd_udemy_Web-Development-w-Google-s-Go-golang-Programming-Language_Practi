package main

import (
	"log"
	"os"
	"text/template"
)

type BevrageStore struct {
	Lager     int
	Ale       int
	Stout     int
	Customers string
}

type Cars struct {
	Sedan      int
	COUPE      int
	SPORTS_CAR int
	Customers  string
}
type Item struct {
	Be []BevrageStore
	Ca []Cars
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	bev := []BevrageStore{
		{230, 45, 23, "Cheif Secretary"},
		{230, 45, 23, "Cheif UFO"},
		{230, 45, 23, "Cheif Vice Pre"},
		{230, 45, 23, "Cheif Presedent"},
		{230, 45, 23, "Cheif Secretary"},
		{230, 45, 23, "Cheif Secretary General"},
		{230, 45, 23, "Cheif Secretary Head"},
		{230, 45, 23, "Deputy Cheif Secretary "},
		{230, 45, 23, "Deputy Cheif Secretary General"},
	}
	ca := []Cars{
		{25551, 2772, 15, "American"},
		{25551, 2772, 15, "American"},
		{251, 252, 15, "American"},
		{251, 22, 15, "American"},
		{2551, 252, 15, "American"},
		{251, 22, 15, "American"},
		{201, 252, 15, "American"},
		{201, 22, 15, "American"},
		{210, 22, 15, "American"},
		{26, 262, 15, "American"},
		{26, 262, 15, "American"},
		{22, 262, 15, "American"},
		{24, 272, 15, "American"},
		{26, 200, 15, "American"},
		{25, 2882, 15, "American"},
	}
	Ite := Item{
		Be: bev,
		Ca: ca,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Ite)
	if err != nil {
		log.Fatal(err)
	}

}
