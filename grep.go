package main

import (
	"bufio"
	"os"
	"strings"
)

// Grep command: grep PATTERN files [...]

type Grep struct {
	pattern string
	files   []string
}

func grep(pattern, file string) []string {
	res := make([]string, 0)
	if exists(file) {
		newFile, _ := os.Open(file)
		defer newFile.Close()
		lines, _, _, _ := wc(file)
		scanner := bufio.NewScanner(newFile)
		for i := 0; i < lines; i++ {
			scanner.Scan()
			if strings.Contains(scanner.Text(), pattern) {
				res = append(res, scanner.Text())
			}
		}
	}
	return res
}

func (g Grep) executeCommand() []string {
	var res []string
	for i, _ := range g.files {
		resFromGrep := grep(g.pattern, g.files[i])
		res = append(res, resFromGrep...)
	}
	return res
}

func (g Grep) readCommand(input string) (string, []string) {
	g.files = divideFiles(input)
	g.pattern = g.files[1]
	g.files = g.files[2:]
	return g.pattern, g.files
}

func (g Grep) startCommand(s string) {
	g.pattern, g.files = g.readCommand(s)
	file := g.executeCommand()
	printSlice(file)
}
