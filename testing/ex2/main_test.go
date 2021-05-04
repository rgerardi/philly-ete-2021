package main

import (
	"bytes"
	"testing"
)

//
//
func Test_Save(t *testing.T) {
	buf := &bytes.Buffer{} // HL
	testdata := "Test data"

	if err := save(buf, []byte(testdata)); err != nil {
		t.Fatal("Error saving data")
	}

	if buf.String() != testdata {
		t.Errorf("Expected %s, got %s\n", testdata, buf.String())
	}
}
