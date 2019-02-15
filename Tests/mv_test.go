package main

import (
	"testing"
)

func TestReadComMv(t *testing.T) {
	var m Mv
	test := "mv 1.txt 2.txt\n"
	_, m.files = m.readCommand(test)
	var expFiles = []string{"1.txt", "2.txt"}
	if len(m.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range m.files {
		if m.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComMv2(t *testing.T) {
	var m Mv
	test := "mv 1.txt 2.txt 3.txt"
	_, m.files = m.readCommand(test)
	var expFiles []string
	if len(m.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range m.files {
		if m.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComMv3(t *testing.T) {
	var m Mv
	test := "mv error.txt"
	_, m.files = m.readCommand(test)
	var expFiles []string
	if len(m.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range m.files {
		if m.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}

func TestReadComMv4(t *testing.T) {
	var m Mv
	test := "mv "
	_, m.files = m.readCommand(test)
	var expFiles []string
	if len(m.files) != len(expFiles) {
		t.Error("Expected files ", expFiles)
	}
	for i, _ := range m.files {
		if m.files[i] != expFiles[i] {
			t.Error("Expected ", expFiles[i])
		}
	}
}
