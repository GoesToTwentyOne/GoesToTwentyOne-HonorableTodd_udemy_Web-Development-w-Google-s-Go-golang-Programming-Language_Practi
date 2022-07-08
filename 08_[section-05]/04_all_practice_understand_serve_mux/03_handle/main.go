package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int
type catdog int

func (c catdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello from catdog</h1>")
}
func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello from hotdog</h1>")
}
func main() {
	var h hotdog
	var c catdog
	http.Handle("/cat/", c)
	http.Handle("/dog/", h)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}

}
