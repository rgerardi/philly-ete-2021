package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	f, _ := ioutil.TempFile("/tmp", "data")
	defer f.Close()

	save(f, data)
}

func save(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}
