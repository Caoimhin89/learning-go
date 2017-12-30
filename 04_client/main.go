package main

import (
	"fmt"
	"net"
)

func main() {
	// establish network connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// write to connection
	fmt.Fprintf(conn, "I dialed you.")
}