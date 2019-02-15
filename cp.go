package main

import (
	"io"
	"os"
)

// Command copy: copy Source Destination, where Destination is a file
type Copy struct {
	files []string
}

func (c Copy) executeCommand() []string {
	var res []string
	if exists(c.files[0]) {
		src, err := os.Open(c.files[0])
		defer src.Close()
		lines, _, _, _ := wc(c.files[0])
		res = head(lines, c.files[0])
		dst, err := os.Create(c.files[1])
		if err != nil {
			panic(err)
		}
		defer dst.Close()
		io.Copy(dst, src)
	}
	return res
}

func (c Copy) readCommand(input string) (flag string, f []string) {
	flag = ""
	if countWords(input) != 3 {
		printErrorCp(input[3 : len(input)-1])
		return flag, c.files
	}
	remain := input[3:]
	c.files = divideFiles(remain)
	return flag, c.files
}

func (c Copy) startCommand(s string) {
	_, c.files = c.readCommand(s)
	_ = c.executeCommand()
}
