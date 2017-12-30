package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// net.Listen asks for type of network and port number to listen on
	// returns a Listener and an error
	// a Listener has 3 methods
	// 1) Accept() (Conn, error) - waits for and returns next connection to Listener
	// 2) Close() error - closes the Listener, blocked Accept ops will be unblocked and return errors
	// 3) Addr() Addr - returns the listener's network address
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Connection has a Reader and a Writer
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// write to connection
		// conn can be passed in as a Writer because it implements the Writer interface
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		// close the connection
		conn.Close()
	}
}
