package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

//START OMIT
func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Listening on port 3000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	msg, err := ioutil.ReadAll(conn)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s\n", msg)
	conn.Close()
}

//END OMIT
