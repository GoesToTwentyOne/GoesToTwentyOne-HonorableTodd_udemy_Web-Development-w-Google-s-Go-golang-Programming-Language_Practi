package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "This is Nihad Hossain"
	tpl := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>Rudimentary concat</title>
	</head>
	<body>
	<h1>` + name + `</h1>

	</body>
	</html>`
	fmt.Println(tpl)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(nf, strings.NewReader(tpl))
}
