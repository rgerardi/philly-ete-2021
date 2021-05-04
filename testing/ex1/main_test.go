package main

import (
	"io/ioutil"
	"os"
	"testing"
)

//
func Test_Save(t *testing.T) {
	testdata := "Test data"
	f, err := ioutil.TempFile("/tmp", "data")
	if err != nil {
		t.Fatal("Error creating test file")
	}
	defer os.Remove(f.Name())

	if err := save(f, []byte(testdata)); err != nil { // HL
		t.Fatal("Error saving data")
	}
	f.Close()

	data, _ := ioutil.ReadFile(f.Name())
	if string(data) != testdata { // HL
		t.Errorf("Expected %s, got %s\n", testdata, string(data))
	}
}
