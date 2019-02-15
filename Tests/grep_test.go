package main

import (
	"io"
	"os"
	"testing"
)

func TestReadComGrep(t *testing.T) {
	var g Grep
	test := "grep a 1.txt 2.txt"
	g.pattern, g.files = g.readCommand(test)
	expPattern := "a"
	var expFiles = []string{"1.txt", "2.txt"}
	if g.pattern != expPattern {
		t.Error("Expected pattern: ", expPattern)
	}
	if len(g.files) != len(expFiles) {
		t.Error("Expected files: ", expFiles)
	}
	for i, _ := range g.files {
		if g.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestExecComGrep(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "abc\nbcad\ncde\ndef\nefg\nfgha\nghi\n")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	res := grep("a", "res.txt")
	var expected = []string{"abc", "bcad", "fgha"}
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
