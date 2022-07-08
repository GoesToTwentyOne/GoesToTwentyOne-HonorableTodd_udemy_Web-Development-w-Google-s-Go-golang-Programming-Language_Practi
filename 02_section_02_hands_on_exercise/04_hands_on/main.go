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

type Item struct {
	Grilled                        string
	Tomato                         string
	Chicken                        string
	Flatbread                      string
	Mini                           string
	Starbucks                      string
	Pizza_Hut_Hong_Kong            string
	BBQ                            string
	Frankfurter                    string
	Burger                         string
	Pizza_Hut                      string
	Bakery_and_Bread               string
	Meat_and_Seafood               string
	Pasta_and_Rice                 string
	Oils, Sauces, Salad, Dressings string
	Cereals_Breakfast              string
	Soups_and_Canned               string
	Frozen                         string
	Dairy, Cheese, Eggs            string
}
type menu struct {
	Resturant_Name string
	Item
}

func main() {
	rsfoood := []menu{
		{
			Resturant_Name: "Indian",
			Item: Item{
				Grilled: "Yes",
				BBQ:     "Yes",

				Tomato:              "Yes",
				Chicken:             "Yes",
				Flatbread:           "Yes",
				Mini:                "NO",
				Starbucks:           "Yes",
				Pizza_Hut_Hong_Kong: "Yes",

				Frankfurter:       "NO",
				Burger:            "Yes",
				Pizza_Hut:         "Yes",
				Bakery_and_Bread:  "NO",
				Meat_and_Seafood:  "Yes",
				Pasta_and_Rice:    "NO",
				Dressings:         "Yes",
				Cereals_Breakfast: "NO",
				Soups_and_Canned:  "NO",
				Frozen:            "NO",
				Eggs:              "NO",
			},
		},
		{
			Resturant_Name: "Texina",
			Item: Item{
				Grilled: "Yes",
				BBQ:     "Yes",

				Tomato:              "Yes",
				Chicken:             "Yes",
				Flatbread:           "Yes",
				Mini:                "NO",
				Starbucks:           "Yes",
				Pizza_Hut_Hong_Kong: "Yes",

				Frankfurter:       "NO",
				Burger:            "Yes",
				Pizza_Hut:         "Yes",
				Bakery_and_Bread:  "NO",
				Meat_and_Seafood:  "Yes",
				Pasta_and_Rice:    "NO",
				Dressings:         "Yes",
				Cereals_Breakfast: "NO",
				Soups_and_Canned:  "NO",
				Frozen:            "NO",
				Eggs:              "NO",
			},
		},
		{
			Resturant_Name: "Bengaliena",
			Item: Item{
				Grilled: "Yes",
				BBQ:     "Yes",

				Tomato:              "Yes",
				Chicken:             "Yes",
				Flatbread:           "Yes",
				Mini:                "NO",
				Starbucks:           "Yes",
				Pizza_Hut_Hong_Kong: "Yes",

				Frankfurter:       "NO",
				Burger:            "Yes",
				Pizza_Hut:         "Yes",
				Bakery_and_Bread:  "NO",
				Meat_and_Seafood:  "Yes",
				Pasta_and_Rice:    "NO",
				Dressings:         "Yes",
				Cereals_Breakfast: "NO",
				Soups_and_Canned:  "NO",
				Frozen:            "NO",
				Eggs:              "NO",
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", rsfoood)
	if err != nil {
		log.Fatal(err)
	}

}
