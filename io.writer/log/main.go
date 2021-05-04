package main

import (
	"fmt"
	"io"
	"net"
)

// START OMIT
func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	client, _ := net.Dial("tcp", ":3000")
	defer client.Close()

	save(LogWriter(client), data) // HL
}

type logWriter struct {
	writer io.Writer
}

func (l logWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("%s\n", p)
	return l.writer.Write(p)
}

func LogWriter(w io.Writer) io.Writer { // HL
	return logWriter{w} // HL
} // HL

// END OMIT

func save(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}
