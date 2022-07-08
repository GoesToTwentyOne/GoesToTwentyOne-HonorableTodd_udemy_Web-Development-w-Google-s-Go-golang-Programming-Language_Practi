package main

import (
	"io"
	"log"
	"net/http"
)

func cat(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello from cat</h1>")
}
func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello from dog</h1>")
}

func main() {
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/dog/", dog)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}

}
