package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Nihad", "Golang Expert")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Hello Bangladesh</h1>")

}
func main() {
	var hd hotdog
	http.ListenAndServe(":5050", hd)

}
