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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := ([]byte(ln))
		r := rot13(bs)
		fmt.Fprintf(conn, "%s -- %s", ln, r)

	}
	defer conn.Close()

}
func rot13(bs []byte) []byte {
	r13 := make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13

}
