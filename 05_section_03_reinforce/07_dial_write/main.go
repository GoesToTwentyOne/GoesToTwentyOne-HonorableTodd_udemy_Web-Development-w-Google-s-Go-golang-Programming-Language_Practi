package main

import (
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	io.WriteString(conn, "I dialed you")

}
