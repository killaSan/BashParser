package main

import (
	"os"
)

// Cat command: cat files[...]

type Cat struct {
	files []string
}

func cat(file string) []string {
	var res []string
	if exists(file) {
		src, _ := os.Open(file)
		defer src.Close()
		lines, _, _, _ := wc(file)
		res = head(lines, file)
	}
	return res
}

func (c Cat) executeCommand() []string {
	var res []string
	for i, _ := range c.files {
		res = append(res, cat(c.files[i])...)
	}
	return res

}

func (c Cat) readCommand(input string) (string, []string) {
	flag := ""
	remain := input[4:]
	c.files = divideFiles(remain)
	return flag, c.files
}

func (c Cat) startCommand(s string) {
	_, c.files = c.readCommand(s)
	res := c.executeCommand()
	printSlice(res)
}
