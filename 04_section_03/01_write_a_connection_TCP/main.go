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
		io.WriteString(conn, "Hello from TCP\n")
		fmt.Fprintln(conn, "Hello from TCP2")
		fmt.Fprintf(conn, "%v", "Hello from TCP3\n")
		conn.Close()
	}

}
