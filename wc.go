package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//WC command: wc [option: w,l,c] files[...]

type Wc struct {
	flag  string
	files []string
}

func wc(file string) (int, int, int, string) {
	var lines, words, char int
	if exists(file) {
		f, _ := os.Open(file)
		defer f.Close()
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			_ = scanner.Text()
			words += 1
		}
		f.Seek(0, 0)
		stat, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		secScan := bufio.NewScanner(f)
		for secScan.Scan() {
			lines += 1
		}
		c := stat.Size()
		char = int(c)
	}

	return lines, words, char, file
}

func (w Wc) executeCommand() []string {
	var res []string
	var lines, words, char int
	var file string
	var resFromWC []string
	for i, _ := range w.files {
		lines, words, char, file = wc(w.files[i])
		if w.flag == "" || (containsSymb(w.flag, 'w') && containsSymb(w.flag, 'c') && containsSymb(w.flag, 'l')) {
			resFromWC = append(resFromWC, strconv.Itoa(lines), strconv.Itoa(words), strconv.Itoa(char), file)
		} else if len(w.flag) == 2 && containsSymb(w.flag, 'l') && containsSymb(w.flag, 'w') {
			resFromWC = append(resFromWC, strconv.Itoa(lines), strconv.Itoa(words), file)
		} else if len(w.flag) == 2 && containsSymb(w.flag, 'l') && containsSymb(w.flag, 'c') {
			resFromWC = append(resFromWC, strconv.Itoa(lines), strconv.Itoa(char), file)
		} else if len(w.flag) == 2 && containsSymb(w.flag, 'w') && containsSymb(w.flag, 'c') {
			resFromWC = append(resFromWC, strconv.Itoa(words), strconv.Itoa(char), file)
		} else if len(w.flag) == 1 && containsSymb(w.flag, 'l') {
			resFromWC = append(resFromWC, strconv.Itoa(lines), file)
		} else if len(w.flag) == 1 && containsSymb(w.flag, 'w') {
			resFromWC = append(resFromWC, strconv.Itoa(words), file)
		} else if len(w.flag) == 1 && containsSymb(w.flag, 'c') {
			resFromWC = append(resFromWC, strconv.Itoa(char), file)
		} else {
			continue
		}
		res = append(res, resFromWC...)
		resFromWC = nil
	}
	return res
}

func (w Wc) readCommand(input string) (string, []string) {
	remain := input[3:]
	if remain[0] == '-' {
		w.flag = remain[1:strings.IndexByte(remain, ' ')]
		remain = remain[len(w.flag)+2:]
	}
	w.files = divideFiles(remain)
	return w.flag, w.files
}

func (w Wc) startCommand(s string) {
	w.flag, w.files = w.readCommand(s)
	res := w.executeCommand()
	printRow(res, len(w.flag)+1)
}
