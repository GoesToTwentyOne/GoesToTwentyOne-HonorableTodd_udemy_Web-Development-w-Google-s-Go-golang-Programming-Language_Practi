package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	nameMy := "Nihad"

	tpl := fmt.Sprint(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>simple concatenation</title>
		</head>
		<body>
			<h1>` + nameMy + `</h1>
		</body>
	</html>
	`)
	nf, err := os.Create("tpl.html")
	if err != nil {
		log.Fatal(err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))

}
