package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	client, _ := net.Dial("tcp", ":3000") // HL
	defer client.Close()                  // HL

	save(client, data)
}

func save(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}
