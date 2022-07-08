package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(conn, "I Love you so I dial you")
	conn.Close()
}
