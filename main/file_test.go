package main

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := os.ReadFile("../testdata/test.data")
	if err != nil {
		t.Fatal("File not read successfully", err)
	}

	if string(f) != "hello world" {
		t.Fatal("hello world file is compromised")
	}
}
