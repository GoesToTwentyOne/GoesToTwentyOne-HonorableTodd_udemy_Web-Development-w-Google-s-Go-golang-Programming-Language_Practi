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
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	<h1>My ` + name + `</h1>
	</body>
	</html>

	`)
	nf, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
