package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

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
	request(conn)
	defer conn.Close()

}
func request(conn net.Conn) {
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
	uri := strings.Fields(ln)[0]
	if method == "GET" && uri == "/" {
		indexRes(conn)

	}

}
func indexRes(conn net.Conn) {

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(../index.gohtml))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	// fmt.Fprint(conn, tpl.ExecuteTemplate(os.Stdout, "index.gohtml", nil))
	tpl.ExecuteTemplate(conn, "../index.gohtml", nil)

}
