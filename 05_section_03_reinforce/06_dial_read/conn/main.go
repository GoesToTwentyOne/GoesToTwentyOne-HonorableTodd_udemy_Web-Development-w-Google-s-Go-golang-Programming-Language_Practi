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

		fmt.Fprintln(conn, "Hello testing")
		fmt.Fprintf(conn, "%v", "testing is good for digging more!")

		conn.Close()

	}
}
