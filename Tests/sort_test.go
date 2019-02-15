package main

import (
	"testing"
)

func TestReadComSort(t *testing.T) {
	var s Sort
	test := "sort -r 1.txt 2.txt 3.txt"
	s.flag, s.files = s.readCommand(test)
	expFlag := "r"
	var expFiles = []string{"1.txt", "2.txt", "3.txt"}
	if expFlag != s.flag {
		t.Error("Expected flag: ", expFlag)
	}
	if len(s.files) != len(expFiles) {
		t.Error("Expected files: ", expFiles)
	}
	for i, _ := range s.files {
		if s.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestReadComSort2(t *testing.T) {
	var s Sort
	test := "sort 1.txt 2.txt 3.txt"
	s.flag, s.files = s.readCommand(test)
	expFlag := ""
	var expFiles = []string{"1.txt", "2.txt", "3.txt"}
	if expFlag != s.flag {
		t.Error("Expected flag: ", expFlag)
	}
	if len(s.files) != len(expFiles) {
		t.Error("Expected files: ", expFiles)
	}
	for i, _ := range s.files {
		if s.files[i] != expFiles[i] {
			t.Error("Expected file: ", expFiles[i])
		}
	}
}

func TestMergeSort(t *testing.T) {
	var expected = []string{"a", "b", "c", "d", "e", "f", "g"}
	res := mergeSort(expected, func(a, b string) bool { return a < b })
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected line: ", expected[i])
		}
	}
}

func TestMergeSort2(t *testing.T) {
	var expected = []string{"g", "f", "e", "d", "c", "b", "a"}
	res := mergeSort([]string{"a", "b", "c", "d", "e", "f", "g"}, func(a, b string) bool { return a > b })
	if len(res) != len(expected) {
		t.Error("Different result expected ", expected)
	}
	for i, _ := range res {
		if res[i] != expected[i] {
			t.Error("Expected line: ", expected[i])
		}
	}
}
