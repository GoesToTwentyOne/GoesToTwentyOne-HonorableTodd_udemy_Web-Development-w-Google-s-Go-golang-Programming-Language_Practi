package main

import "fmt"

func main() {
	nameMy := "Nihad Hossain"
	tpl := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>simple concatenation</title>
		</head>
		<body>
			<h1>` + nameMy + `</h1>
		</body>
	</html>
	`
	fmt.Println(tpl)

}
