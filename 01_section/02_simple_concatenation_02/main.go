package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Nihad Hossain"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>simple concatenation</title>
	</head>
	<body>
	<h1>My ` + name + `</h1>
	</body>
	</html>
	`)
	// fmt.Println(str)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(nf, strings.NewReader(str))

}
