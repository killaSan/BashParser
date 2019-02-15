package main

import (
	"io"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestHead(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "a\nb\nc\nd\ne\nf\ng\n")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	res := head(3, "res.txt")
	var expected = []string{"a", "b", "c"}
	err = os.Remove("res.txt")
	if err != nil {
		panic(err)
	}
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Error!")
		}
	}
}

func TestReadComHead(t *testing.T) {
	var h Head
	test := "head file1.txt file2.txt"
	h.flag, h.files = h.readCommand(test)
	expFlag := "10"
	var expFiles = []string{"file1.txt", "file2.txt"}
	if h.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(h.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range h.files {
		if h.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComHead2(t *testing.T) {
	var h Head
	test := "head -3 1"
	h.flag, h.files = h.readCommand(test)
	expFlag := "3"
	var expFiles = []string{"1"}
	if h.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(h.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range h.files {
		if h.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComHead3(t *testing.T) {
	var h Head
	test := "head -40 file1.txt file2.txt fileN.go"
	h.flag, h.files = h.readCommand(test)
	expFlag := "40"
	var expFiles = []string{"file1.txt", "file2.txt", "fileN.go"}
	if h.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(h.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range h.files {
		if h.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}
