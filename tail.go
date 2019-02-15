package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func Tail: tail [-n] files [...]

type Tail struct {
	flag  string
	files []string
}

func tail(n int, file string) []string {
	res := make([]string, 0)
	if exists(file) {
		lines, _, _, _ := wc(file)
		if n >= lines {
			res = head(lines, file)
		} else {
			skipLines := lines - n
			var i int
			newFile, _ := os.Open(file)
			defer newFile.Close()
			scanner := bufio.NewScanner(newFile)
			for i = 0; i < skipLines; i++ {
				scanner.Scan()
			}
			for i = skipLines; i <= lines; i++ {
				if scanner.Scan() {
					res = append(res, scanner.Text())
				}
			}
		}
	}
	return res
}

func (t Tail) executeCommand() []string {
	var res []string
	flag, err := strconv.Atoi(t.flag)
	if err != nil {
		message := fmt.Sprintf("head: invalid option --'%v'", t.flag)
		fmt.Println(message)
		fmt.Println()
		return res
	}
	for i, _ := range t.files {
		resFromTail := tail(flag, t.files[i])
		res = append(res, resFromTail...)
	}
	return res
}

func (t Tail) readCommand(input string) (string, []string) {
	var indx int = 5
	var partition string
	if input[:indx] == "tail " {
		if input[indx] == '-' {
			partition = input[indx:]
			t.flag = partition[1:strings.IndexByte(partition, ' ')]
			indx += len(t.flag) + 2
		} else {
			t.flag = "10" // extract last 10 lines from a file
		}
		partition = input[indx:]
		t.files = divideFiles(partition)
	}
	return t.flag, t.files
}

func (t Tail) startCommand(s string) {
	t.flag, t.files = t.readCommand(s)
	res := t.executeCommand()
	printSlice(res)
}
