package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// here we dail the address from the tcp network
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln("Error reading the connection : ", err)
	}

	fmt.Println(string(bs))
}
