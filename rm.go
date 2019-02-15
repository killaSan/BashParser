package main

import (
	"os"
)

// rm command rm files [...]

type Rm struct {
	files []string
}

func rm(file string) {
	if exists(file) {
		_ = os.Remove(file)
	}
	return
}

func (r Rm) executeCommand() []string {
	var empty []string
	for i, _ := range r.files {
		rm(r.files[i])
	}
	return empty
}

func (r Rm) readCommand(input string) (string, []string) {
	flag := ""
	r.files = divideFiles(input)
	r.files = r.files[1:]
	return flag, r.files
}

func (r Rm) startCommand(s string) {
	_, r.files = r.readCommand(s)
	_ = r.executeCommand()
}
