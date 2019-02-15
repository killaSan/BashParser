package main

import (
	"os"
	"testing"
)

func TestReadComRm(t *testing.T) {
	var r Rm
	test := "rm 1 2 3 4"
	_, r.files = r.readCommand(test)
	var expFiles = []string{"1", "2", "3", "4"}
	if len(r.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range r.files {
		if r.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestRmComm(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	rm("res.txt")
	if exists("res.txt") {
		t.Error("Rm didn't remove file: res.txt")
	}
}
