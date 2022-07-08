package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/cat/":
		io.WriteString(w, "<h1>Hello Cat</h1>")
	case "/dog/":
		io.WriteString(w, "<h1>Hello Dog</h1>")
	}
}
func main() {
	var h hotdog
	err := http.ListenAndServe(":5050", h)
	if err != nil {
		log.Fatal(err)
	}

}
