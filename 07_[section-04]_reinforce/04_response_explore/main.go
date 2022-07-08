package main

import (
	"fmt"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Nihad", "Hossain")
	w.Header().Set("Content-Type", "Text/html")
	fmt.Fprintln(w, "<h1>Hello</h1>")
}
func main() {
	var h hotdog
	err := http.ListenAndServe(":5050", h)
	if err != nil {
		log.Fatal(err)
	}

}
