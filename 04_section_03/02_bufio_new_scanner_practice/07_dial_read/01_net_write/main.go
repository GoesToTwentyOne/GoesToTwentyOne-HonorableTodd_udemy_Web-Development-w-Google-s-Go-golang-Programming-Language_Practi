package main

import (
	"fmt"
	"io"
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
		io.WriteString(conn, "Hi I'm NIhad Hossain")
		fmt.Fprintln(conn, "Hi I'm From Bangladesh")
		fmt.Fprintf(conn, "Hi I'm Junior Engineer at %s \n", "UN")
		conn.Close()

	}
}
