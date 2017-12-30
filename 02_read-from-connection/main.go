package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Accept incoming connections
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// handle more than one connection if necessary
		go handle(conn)
	}
}

// handle function will loop infinitely, connection will not be closed
func handle(conn net.Conn) {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		defer conn.Close()
}
