package main

import "fmt"

func main() {
	name := "Nihad Hossain"
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head>
	<title>simple concatenation</title>
	</head>
	<body>
	<h1>My ` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tmpl)

}
