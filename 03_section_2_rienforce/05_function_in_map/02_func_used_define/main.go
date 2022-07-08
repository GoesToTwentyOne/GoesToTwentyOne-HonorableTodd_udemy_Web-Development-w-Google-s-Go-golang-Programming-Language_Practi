package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fn).ParseFiles("tpl.gohtml"))
}
func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}
func sumcEvenSlice(s []int) int {
	sum := 0
	for _, v := range s {
		if v%2 == 0 {
			sum += v

		}
	}
	return sum
}
func sumcOddSlice(s []int) int {
	sum := 0
	for _, v := range s {
		if v%2 != 0 {
			sum += v

		}
	}
	return sum
}

var fn = template.FuncMap{
	"sxs":  sumSlice,
	"sxso": sumcOddSlice,
	"sxse": sumcEvenSlice,
}

type Data struct {
	Num []int
}

func main() {
	n := []Data{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{91, 92, 93, 94, 95, 96, 97, 98, 99}},
		{[]int{60, 61, 62, 63, 64, 65, 66, 67, 68, 69}},
		{[]int{70, 71, 72, 73, 74, 75, 76, 77, 78, 79}},
		{[]int{80, 81, 82, 83, 84, 85, 86, 87, 88, 89}},
		{[]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
		{[]int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}},
		{[]int{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", n)
	if err != nil {
		log.Fatal(err)
	}

}
