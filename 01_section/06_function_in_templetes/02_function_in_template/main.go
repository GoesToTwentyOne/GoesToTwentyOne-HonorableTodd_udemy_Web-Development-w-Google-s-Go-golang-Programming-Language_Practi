package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("tpl.gohtml"))

}

var ft = template.FuncMap{
	"US":  strings.ToUpper,
	"FTS": firsThreeCap,
}

func firsThreeCap(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return strings.ToUpper(s)

}

type WesternPantsGirls struct {
	CowgirlLoved  string
	PartywearLove string
	SleeveLove    string
}
type WesternPantsBoys struct {
	CowboyLoved string
	HatLove     string
	ShirtLove   string
}

func main() {
	weG := []WesternPantsGirls{
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
	}
	weB := []WesternPantsBoys{
		{"Banglaes", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
		{"texas", "califonia", "silicon_velly"},
		{"India", "West Bengal", "Maharastra"},
		{"China", "Whuan", "Whuan Institute"},
		{"England", "London", "Bangladesh community"},
	}
	wearParsing := struct {
		WearGirl []WesternPantsGirls
		WearBoy  []WesternPantsBoys
	}{
		WearGirl: weG,
		WearBoy:  weB,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", wearParsing)
	if err != nil {
		log.Fatal(err)
	}

}
