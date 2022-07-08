package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type foodDishes struct {
	Name, Price string
}
type item struct {
	ResName                   string
	Break_fast, Lunch, Dinner []foodDishes
}

func main() {
	ritem := []item{
		{
			ResName: "Hotlite",
			Break_fast: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
			Lunch: []foodDishes{
				{Name: "Calas", Price: "$5"},
				{Name: "Charleston red rice", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Dirty rice", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Spanish rices", Price: "$16"},
			},
			Dinner: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Chicken bog", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
		},
		{
			ResName: "Daylight",
			Break_fast: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
			Lunch: []foodDishes{
				{Name: "Calas", Price: "$5"},
				{Name: "Charleston red rice", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Dirty rice", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Spanish rices", Price: "$16"},
			},
			Dinner: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Chicken bog", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
		},
		{
			ResName: "NightLight",
			Break_fast: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
			Lunch: []foodDishes{
				{Name: "Calas", Price: "$5"},
				{Name: "Charleston red rice", Price: "$7"},
				{Name: "Eggs Benedict", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Dirty rice", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Spanish rices", Price: "$16"},
			},
			Dinner: []foodDishes{
				{Name: "Pancakes and Maple Syrup", Price: "$5"},
				{Name: "Toasted English Muffin", Price: "$7"},
				{Name: "Chicken bog", Price: "$6"},
				{Name: "Biscuits and Gravy", Price: "$8"},
				{Name: "Cinnamon Rolls", Price: "$12"},
				{Name: "Belgian Style Waffles", Price: "$23"},
				{Name: "Breakfast Burrito", Price: "$14"},
				{Name: "French Toast", Price: "$21"},
				{Name: " Bacon and Eggs", Price: "$16"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ritem)
	if err != nil {
		log.Fatal(err)
	}

}
