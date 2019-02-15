package main

import (
	"testing"
)

func TestReadComCp(t *testing.T) {
	var c Copy
	test := "cp 1.txt 2.txt"
	_, c.files = c.readCommand(test)
	var expFiles = []string{"1.txt", "2.txt"}
	if len(c.files) != len(expFiles) {
		t.Error("Expected", expFiles)
	}
	for i, _ := range c.files {
		if c.files[i] != expFiles[i] {
			t.Error("Expected file ", expFiles[i])
		}
	}
}

func TestReadComCp2(t *testing.T) {
	var c Copy
	test := "cp 1.txt 2.txt err.txt"
	_, c.files = c.readCommand(test)
	var expFiles []string
	if len(c.files) != len(expFiles) {
		t.Error("Expected", expFiles)
	}
	for i, _ := range c.files {
		if c.files[i] != expFiles[i] {
			t.Error("Expected file ", expFiles[i])
		}
	}
}
