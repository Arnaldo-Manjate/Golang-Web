package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Run this program with this command telnet loaclhost 8080
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error Listening on port 8080 :", err)
	}
	defer listener.Close()

	// after returning the listener listen forever for a connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error Connecting to port 8080 :", err)
		}
		defer conn.Close()

		_, err = io.WriteString(conn, fmt.Sprintf("Hello from Server \n %v", time.Now()))
		if err != nil {
			log.Fatalln("Error writing to the Connection", err)
		}
	}
}
