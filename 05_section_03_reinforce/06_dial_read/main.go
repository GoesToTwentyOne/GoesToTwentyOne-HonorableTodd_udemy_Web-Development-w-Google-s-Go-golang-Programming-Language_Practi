package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))

}
