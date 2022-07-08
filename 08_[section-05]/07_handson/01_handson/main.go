package main

import (
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Form Home</h1>")
}
func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Form Dog</h1>")
}
func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>I'm Nihad Hossain</h1>")
}
func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
