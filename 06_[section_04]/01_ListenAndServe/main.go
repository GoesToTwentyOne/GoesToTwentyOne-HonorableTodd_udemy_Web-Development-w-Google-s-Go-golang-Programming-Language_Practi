package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>Hello world</strong><br>

	</body></html>`)

}

func main() {
	var dog hotdog
	http.ListenAndServe(":5050", dog)
}
