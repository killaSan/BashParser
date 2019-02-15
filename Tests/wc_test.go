package main

import (
	"io"
	"os"
	"testing"
)

func TestReadComWc(t *testing.T) {
	var w Wc
	test := "wc 1 2 3"
	w.flag, w.files = w.readCommand(test)
	expFlag := ""
	var expFiles = []string{"1", "2", "3"}
	if w.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(w.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range w.files {
		if w.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestReadComWc2(t *testing.T) {
	var w Wc
	test := "wc -lw 1 2 3"
	w.flag, w.files = w.readCommand(test)
	expFlag := "lw"
	var expFiles = []string{"1", "2", "3"}
	if w.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(w.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range w.files {
		if w.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestReadComWc3(t *testing.T) {
	var w Wc
	test := "wc -c 1 2 3"
	w.flag, w.files = w.readCommand(test)
	expFlag := "c"
	var expFiles = []string{"1", "2", "3"}
	if w.flag != expFlag {
		t.Error("Expected flag ", expFlag)
	}
	if len(w.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range w.files {
		if w.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestWc(t *testing.T) {
	f, err := os.Create("res.txt")
	check(err)
	defer f.Close()
	_, err = io.WriteString(f, "a\nb\nc\nd\ne\nf\ng h\n")
	if err != nil {
		panic(err)
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		panic(err)
	}
	lines, words, char, file := wc("res.txt")
	expLines := 7
	expWords := 8
	expChar := 16
	expFile := "res.txt"
	err = os.Remove("res.txt")
	if err != nil {
		panic(err)
	}
	if lines != expLines {
		t.Error("Expected lines: ", expLines)
	}
	if words != expWords {
		t.Error("Expected words: ", expWords)
	}
	if char != expChar {
		t.Error("Expected chars: ", expChar)
	}
	if file != expFile {
		t.Error("Expected file: ", expFile)
	}
}
