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

type Meal struct {
	Meal_name string
}
type Item struct {
	Meal
	Food_Name    string
	Descrip      string
	Price_dollar float64
	Place        string
}

func main() {
	it := []Item{
		{
			Meal: Meal{
				Meal_name: "Dinner",
			},
			Food_Name:    "Oatmeal",
			Descrip:      "Yume Yume",
			Price_dollar: 5.5,
			Place:        "Room",
		},
		{
			Meal: Meal{
				Meal_name: "Lunch",
			},
			Food_Name:    "Hamburger",
			Descrip:      "Delicious and good food",
			Price_dollar: 5.5,
			Place:        "Canteen",
		},
		{
			Meal: Meal{
				Meal_name: "Dinner",
			},
			Food_Name:    "Pasta",
			Descrip:      "take test and tyr again incredible!",
			Price_dollar: 5.5,
			Place:        "Room",
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", it)
	if err != nil {
		log.Fatal(err)
	}

}
