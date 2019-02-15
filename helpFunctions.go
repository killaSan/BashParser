package main

import (
	"fmt"
	"os"
	"strings"
)

// exists returns true if a file exists, otherwise - false
func exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		res := fmt.Sprintf("Cannot open %v : No such file or directory", name)
		fmt.Println(res)
	}
	return (err == nil)
}

// countWords returns the number of files passed in the main function
func countWords(name string) int {
	if len(name) == 0 {
		return 0
	}
	words := 1
	for i := 0; i < len(name)-1; i++ {
		if name[i] == ' ' && name[i+1] == '\n' {
			continue
		} else if name[i] == ' ' {
			words += 1
		}
	}
	return words
}

// printErrorHead prints an error if the symbol passed as a flag is not a number
func printErrorHead(name, wrongNum string) {
	res := fmt.Sprintf("head: invalid option --'%v'", wrongNum)
	fmt.Println(res)
	fmt.Println()
}

// printErrorMv prints an error if only one file is passed to mv command
func printErrorMv(name string) {
	res := fmt.Sprintf("mv: missing destination file operand after '%v'", name)
	fmt.Println(res)
}

// printErrorCp prints an error if only one file is passed to cp command
func printErrorCp(name string) {
	res := fmt.Sprintf("cp: missing destination file operand after '%v'", name)
	fmt.Println(res)
}

// printErrorCut prints an error if numbers are not passed to position in the cut command
func printErrorCut(name string) {
	res := fmt.Sprintf("cut: invalid byte/character position'%v'", name)
	fmt.Println(res)
}

// printSlice prints the result of an operation
func printSlice(text []string) {
	for _, t := range text {
		fmt.Println(t)
	}
}

//
func printRow(text []string, c int) {
	var cnt int
	if c == 1 {
		c = 4
	}
	for _, t := range text {
		if cnt == c {
			fmt.Println()
			cnt = 0
		}
		fmt.Print(t)
		fmt.Print(" ")
		cnt += 1
	}
}

// divideFiles returns the files passed to a command
func divideFiles(remain string) []string {
	files := make([]string, 0)
	words := countWords(remain)
	for i := 0; i < words-1; i++ {
		file := remain[0:strings.IndexByte(remain, ' ')]
		files = append(files, file)
		remain = remain[len(file)+1:]
	}
	if remain[len(remain)-1] == '\n' {
		files = append(files, remain[:len(remain)-1])
	} else {
		files = append(files, remain[:len(remain)])
	}
	return files
}

// ContainsSymb returns true if a string contains res(char) in it
func containsSymb(s string, res byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == res {
			return true
		}
	}
	return false
}

// removeTabs removes all tabs at the beginnig of a string for sort command
func removeTabs(s string) string {
	if len(s) >= 1 {
		for s[0] == '\t' {
			s = s[1:]
		}
	}
	return s
}
