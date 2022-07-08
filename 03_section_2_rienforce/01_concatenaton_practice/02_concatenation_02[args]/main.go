package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
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
