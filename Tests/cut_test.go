package main

import (
	"io"
	"os"
	"testing"
)

func TestExecComCut(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "abc\nbcd\ncde\ndef\nefg\nfgh\nghi\n")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	res := cut(0, 2, "res.txt")
	var expected = []string{"ab", "bc", "cd", "de", "ef", "fg", "gh"}
	err = os.Remove("res.txt")
	if err != nil {
		panic(err)
	}
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected line: ", res[i])
		}
	}
}

func TestExecComCut2(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "abc\nbcd\ncde\ndef\nefg\nfgh\nghi\n")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	res := cut(5, 10, "res.txt")
	var expected = []string{"", "", "", "", "", "", ""}
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

func TestReadComCut(t *testing.T) {
	var c Cut
	test := "cut -c 1-3 1.txt"
	c.flag, c.files = c.readCommand(test)
	expFlag := "1-3"
	var expfiles = []string{"1.txt"}
	if c.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(c.files) != len(expfiles) {
		t.Error("Expected files ", expfiles)
	}
	for i, _ := range c.files {
		if c.files[i] != expfiles[i] {
			t.Error("Expected file: ", expfiles[i])
		}
	}
}

func TestReadComCut2(t *testing.T) {
	var c Cut
	test := "cut 1-3 1.txt 2.txt"
	c.flag, c.files = c.readCommand(test)
	expFlag := "error"
	var expfiles []string
	if c.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(c.files) != len(expfiles) {
		t.Error("Expected files ", expfiles)
	}
	for i, _ := range c.files {
		if c.files[i] != expfiles[i] {
			t.Error("Expected file: ", expfiles[i])
		}
	}
}
