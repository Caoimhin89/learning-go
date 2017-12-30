package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"time"
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

func handle(conn net.Conn) {
	// set deadline for connection, close after 10 seconds
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection Timeout")
	}

	scanner := bufio.NewScanner(conn)
	// read from the connection
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		// write to the connection
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()
}