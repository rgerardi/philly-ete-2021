package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Writing data")
	data := []byte("Some data to write")

	f, _ := ioutil.TempFile("/tmp", "data")
	defer f.Close()

	save(f, data)
}

func save(f *os.File, data []byte) error {
	_, err := f.Write(data)
	return err
}
