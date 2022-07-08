package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>This is a Heading</h1>
	<p>This is a paragraph.</p>
	
	</body>
	</html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d \r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: html/text\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
