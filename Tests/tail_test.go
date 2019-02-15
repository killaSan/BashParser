package main

import (
	"io"
	"os"
	"testing"
)

func TestTail(t *testing.T) {
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
	res := tail(3, "res.txt")
	var expected = []string{"e", "f", "g"}
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

func TestReadComTail(t *testing.T) {
	var tl Tail
	test := "tail file1.txt file2.txt"
	tl.flag, tl.files = tl.readCommand(test)
	expFlag := "10"
	var expFiles = []string{"file1.txt", "file2.txt"}
	if tl.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(tl.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range tl.files {
		if tl.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComTail2(t *testing.T) {
	var tl Tail
	test := "tail -3 1"
	tl.flag, tl.files = tl.readCommand(test)
	expFlag := "3"
	var expFiles = []string{"1"}
	if tl.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(tl.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range tl.files {
		if tl.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComTail3(t *testing.T) {
	var tl Tail
	test := "tail -40 file1.txt file2.txt fileN.go"
	tl.flag, tl.files = tl.readCommand(test)
	expFlag := "40"
	var expFiles = []string{"file1.txt", "file2.txt", "fileN.go"}
	if tl.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(tl.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range tl.files {
		if tl.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}
