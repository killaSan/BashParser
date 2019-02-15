package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Head command: head [-n] files[...]
type Head struct {
	flag  string
	files []string
}

func head(n int, file string) []string {
	res := make([]string, 0)
	if exists(file) {
		newFile, _ := os.Open(file)
		defer newFile.Close()
		scanner := bufio.NewScanner(newFile)
		for i := 0; i < n; i++ {
			if scanner.Scan() {
				res = append(res, scanner.Text())
			} else {
				break
			}
		}
	}
	return res
}

func (h Head) readCommand(input string) (string, []string) {
	var indx int = 5
	var partition string
	if input[indx] == '-' {
		partition = input[indx:]
		h.flag = partition[1:strings.IndexByte(partition, ' ')]
		indx += len(h.flag) + 2
	} else {
		h.flag = "10" // extract first 10 lines from a file
	}
	partition = input[indx:]
	if len(partition) != 0 {
		h.files = divideFiles(partition)
	}
	return h.flag, h.files
}

func (h Head) executeCommand() []string {
	res := make([]string, 0)
	flag, err := strconv.Atoi(h.flag)
	if err != nil {
		message := fmt.Sprintf("head: invalid option --'%v'", h.flag)
		fmt.Println(message)
		fmt.Println()
		return res
	}
	for i, _ := range h.files {
		resFromHead := head(flag, h.files[i])
		res = append(res, resFromHead...)
	}
	return res
}

func (h Head) startCommand(s string) {
	h.flag, h.files = h.readCommand(s)
	res := h.executeCommand()
	printSlice(res)
}
