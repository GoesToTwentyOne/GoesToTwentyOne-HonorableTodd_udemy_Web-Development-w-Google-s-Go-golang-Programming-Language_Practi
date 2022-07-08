package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int
type catdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hollo from hotdog</h1>")
}
func (c catdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hollo from catdog</h1>")
}

func main() {
	var h hotdog
	var c catdog
	mux := http.NewServeMux()
	mux.Handle("/dog/", h)
	mux.Handle("/cat/", c)
	err := http.ListenAndServe(":5050", mux)
	if err != nil {
		log.Fatal(err)
	}

}
