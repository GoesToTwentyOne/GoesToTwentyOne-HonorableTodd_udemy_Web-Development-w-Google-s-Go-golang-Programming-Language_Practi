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
		bs, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bs))
		conn.Close()
	}
}
