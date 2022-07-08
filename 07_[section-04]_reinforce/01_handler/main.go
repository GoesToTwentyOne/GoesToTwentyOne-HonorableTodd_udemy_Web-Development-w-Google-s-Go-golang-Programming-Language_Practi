package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello from handler")

}

func main() {
	var h hotdog
	http.ListenAndServe(":5050", h)
}
