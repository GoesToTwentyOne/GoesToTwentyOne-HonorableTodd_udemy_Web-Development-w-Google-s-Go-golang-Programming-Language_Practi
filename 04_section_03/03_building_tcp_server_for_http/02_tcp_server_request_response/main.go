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
	//request
	request(conn)
	//response
	response(conn)
	defer conn.Close()

}
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method := strings.Fields(ln)[0]
			fmt.Println("*******METHOD:", method)
		}
		if ln == "" {
			break
		}
		i++
	}

}
func response(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
	<body>
	
	<h2>Text field</h2>
	<p>The <strong>input type="text"</strong> defines a one-line text input field:</p>
	
	<form action="/action_page.php">
	  <label for="fname">First name:</label><br>
	  <input type="text" id="fname" name="fname"><br>
	  <label for="lname">Last name:</label><br>
	  <input type="text" id="lname" name="lname"><br><br>
	  <input type="submit" value="Submit">
	</form>
	
	<p>Note that the form itself is not visible.</p>
	<p>Also note that the default width of a text field is 20 characters.</p>
	
	</body>
	</html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length:%d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type:text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
