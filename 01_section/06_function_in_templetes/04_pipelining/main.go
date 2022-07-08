package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(multioperation).ParseFiles("tpl.gohtml"))
}
func double(x float64) float64 {
	return x + x
}
func power(x float64) float64 {
	return math.Pow(x, 2)

}
func rootsqrt(x float64) float64 {
	return math.Sqrt(x)
}

var multioperation = template.FuncMap{
	"dou":   double,
	"pow":   power,
	"rsqrt": rootsqrt,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 4.00)
	if err != nil {
		log.Fatal(err)
	}

}
