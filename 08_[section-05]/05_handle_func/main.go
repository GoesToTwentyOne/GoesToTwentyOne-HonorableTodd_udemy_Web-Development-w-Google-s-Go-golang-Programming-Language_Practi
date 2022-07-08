package main

import (
	"io"
	"log"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello from dog</h1>")
}
func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello from cat</h1>")
}
func main() {

	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}

}
