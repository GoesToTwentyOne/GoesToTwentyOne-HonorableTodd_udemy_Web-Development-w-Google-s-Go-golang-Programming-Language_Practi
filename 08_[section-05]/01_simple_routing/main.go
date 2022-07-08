package main

import (
	"io"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		io.WriteString(w, "<h1>Hello World</h1>")
	case "/name":
		io.WriteString(w, "<h1>My name is Nihad Hossain</h1>")

	case "/age":
		io.WriteString(w, "<h1>I'm 21</h1>")

	}
}
func main() {
	var h hotdog
	err := http.ListenAndServe(":5050", h)
	if err != nil {
		log.Fatal(err)
	}

}
