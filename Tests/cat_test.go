package main

import (
	"io"
	"os"
	"testing"
)

func TestReadComCat(t *testing.T) {
	var c Cat
	test := "cat 1.txt 2.txt 3.txt"
	_, c.files = c.readCommand(test)
	var expFiles = []string{"1.txt", "2.txt", "3.txt"}
	if len(c.files) != len(expFiles) {
		t.Error("Expected files", expFiles)
	}
	for i, _ := range c.files {
		if c.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestExecComCat(t *testing.T) {
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
	res := cat("res.txt")
	var expected = []string{"a", "b", "c", "d", "e", "f", "g"}
	err = os.Remove("res.txt")
	if err != nil {
		panic(err)
	}
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected line: ", expected[i])
		}
	}
}

func TestExecComCat2(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	res := cat("res.txt")
	var expected []string
	err = os.Remove("res.txt")
	if err != nil {
		panic(err)
	}
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected line: ", expected[i])
		}
	}
}
