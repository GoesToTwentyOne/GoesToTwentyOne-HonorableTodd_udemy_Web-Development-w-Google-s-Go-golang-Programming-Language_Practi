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
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++

	}
}
func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Printf("******method:%s\r\n", method)
	fmt.Printf("******uri:%s\r\n", uri)
	if method == "GET" && uri == "/" {
		home(conn)
	}
	if method == "GET" && uri == "/contact" {
		contact(conn)
	}
	if method == "POST" && uri == "/contact" {
		contactprocess(conn)
	}
	if method == "GET" && uri == "/news" {
		news(conn)
	}
	if method == "GET" && uri == "/about" {
		about(conn)
	}

}
func home(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<strong>Home</strong><br>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	</ul>
	</body>
	</html>
	`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type:Text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func news(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<strong>News</strong><br>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	</ul>
	</body>
	</html>
	`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type:Text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<strong>About</strong><br>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	</ul>
	</body>
	</html>
	`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type:Text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<strong>Contact</strong><br>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	  	
	<h2>Text input fields</h2>
	
	<form method="POST" action="/contact">
	  <label for="fname">First name:</label><br>
	  <input type="text" id="fname" name="fname" value="John"><br>
	  <label for="lname">Last name:</label><br>
	  <input type="text" id="lname" name="lname" value="Doe">
	  <input type="submit" value="Submit">
	</form>
	</ul>
	</body>
	</html>
	`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type:Text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)

}
func contactprocess(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	<strong>Contact</strong><br>
	
	<ul>
	  <li><a href="/">Home</a></li>
	  <li><a href="/news">News</a></li>
	  <li><a href="/contact">Contact</a></li>
	  <li><a href="/about">About</a></li>
	</ul>
	</body>
	</html>
	`
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type:Text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprint(conn, body)

}
