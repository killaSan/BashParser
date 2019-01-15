package main

import (
	"bufio"
	"fmt"
	"github.com/killaSan/helperFunctions"
	"os"
	"strconv"
	"strings"
)

type Command interface {
	readCommand(string) (string, []string)
	executeCommand()
}

// Structure of head command
// flag is the number of lines that are going to be printed
// files are the files that are going to be passed to head
type Head struct {
	input string
	flag  string
	files []string
}

func (h Head) readCommand(input string) (string, []string) {
	var indx int = 5
	var partition string
	if input[:indx] == "head " {
		if input[indx] == '-' {
			partition = input[indx:]
			h.flag = partition[1:strings.IndexByte(partition, ' ')]
			indx += len(h.flag) + 2
		} else {
			h.flag = "10" // extract first 10 lines from a file
		}
		partition = input[indx:]
		f := helper.CreateHelper(partition)
		words := f.CountWords()
		for i := 0; i < words-1; i++ {
			file := partition[0:strings.IndexByte(partition, ' ')]
			h.files = append(h.files, file)
			partition = partition[len(file)+1:]
		}
		h.files = append(h.files, partition[:len(partition)-1])
	}
	return h.flag, h.files
}
func head(n int, file string, cnt int) {
	f := helper.CreateHelper(file)
	if !f.Exists() {
		fmt.Println()
		return
	}
	newFile, _ := os.Open(file)
	defer newFile.Close()
	if cnt > 1 {
		message := fmt.Sprintf("==> %v <==", file)
		fmt.Println(message)
	}
	scanner := bufio.NewScanner(newFile)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		} else {
			break
		}
	}
	fmt.Println()
}

func (h Head) executeCommand() {
	flag, err := strconv.Atoi(h.flag)
	if err != nil {
		res := fmt.Sprintf("head: invalid option --'%v'", h.flag)
		fmt.Println(res)
		fmt.Println()
		return
	}
	for i, _ := range h.files {
		head(flag, h.files[i], len(h.files))
	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command: ")
		input, _ := reader.ReadString('\n')
		if input == "exit\n" {
			break
		} else if len(input) >= 5 && input[:5] == "head " {
			var h Head
			h.flag, h.files = h.readCommand(input)
			h.executeCommand()
		} else {
			fmt.Println("Wrong command! ")
		}
	}

}
