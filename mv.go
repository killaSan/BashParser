package main

import (
	"os"
)

// Move command: mv Source Destination, where Destination is a file
type Mv struct {
	files []string
}

func (m Mv) executeCommand() []string {
	var res []string
	if exists(m.files[0]) {
		lines, _, _, _ := wc(m.files[0]) // m.files[0] = source
		res = head(lines, m.files[0])
		err := os.Rename(m.files[0], m.files[1]) // m.files[1] = destination
		if err != nil {
			panic(err)
		}
	}
	return res
}

func (m Mv) readCommand(input string) (string, []string) {
	flag := ""
	if countWords(input) != 3 {
		//printErrorMv(input[3:len(input) - 1])
		printErrorMv(input)
		return flag, m.files
	}
	remain := input[3:]
	m.files = divideFiles(remain)
	return flag, m.files
}

func (m Mv) startCommand(s string) {
	_, m.files = m.readCommand(s)
	_ = m.executeCommand()
}
