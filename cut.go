package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Cut command: cut -c beg-end files[...]
type Cut struct {
	flag  string
	files []string
}

func cut(beg, end int, file string) []string {
	res := make([]string, 0)
	if exists(file) {
		f, _ := os.Open(file)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				res = append(res, line)
			} else if len(line) < beg && len(line) < end {
				res = append(res, "")
			} else if len(line) < beg {
				res = append(res, "")
			} else if len(line) < end {
				res = append(res, line[beg:])
			} else {
				res = append(res, line[beg:end]) //?
			}
		}
	}
	return res
}

func (c Cut) executeCommand() []string {
	var begin, end string
	var res []string
	if c.flag == "error" {
		return res
	}
	begin = c.flag[0:strings.IndexByte(c.flag, '-')]
	end = c.flag[len(begin)+1:]
	beginPos, err := strconv.Atoi(begin)
	if err != nil {
		printErrorCut(begin)
		return res
	}
	endPos, err := strconv.Atoi(end)
	if err != nil {
		printErrorCut(end)
		return res
	}
	beginPos = beginPos - 1
	endPos = endPos
	if beginPos == -1 {
		fmt.Println("cut: byte/character positions are numbered from 1")
		return res
	}
	if beginPos > endPos {
		fmt.Println("cut: invalid decreasing range")
		return res
	}

	for i, _ := range c.files {
		resFromCut := cut(beginPos, endPos, c.files[i])
		res = append(res, resFromCut...)
	}
	return res
}

func (c Cut) readCommand(input string) (string, []string) {
	if input[:6] != "cut -c" {
		fmt.Println("wrong command!")
		return "error", c.files
	}
	remain := input[7:]
	c.flag = remain[0:strings.IndexByte(remain, ' ')]
	if !containsSymb(c.flag, '-') {
		c.flag = "error"
	}
	remain = remain[len(c.flag)+1:]
	c.files = divideFiles(remain)
	return c.flag, c.files
}

func (c Cut) startCommand(s string) {
	c.flag, c.files = c.readCommand(s)
	file := c.executeCommand()
	printSlice(file)
}
