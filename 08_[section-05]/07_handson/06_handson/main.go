package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}

}
func handle(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			method := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]

			fmt.Println(method)
			fmt.Println(uri)
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}
func mux(conn net.Conn, ln string) {
	defer conn.Close()
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	switch {
	case method == "GET" && uri == "/":
		index(conn)
	case method == "GET" && uri == "/cat/":
		cat(conn)
	case method == "GET" && uri == "/dog/":
		dog(conn)

	}

}
func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>Index Home</h1>

	
	</body>
	</html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func cat(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>Cat</h1>

	
	</body>
	</html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func dog(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>Dog</h1>

	
	</body>
	</html>`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)
}
