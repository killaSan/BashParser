package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var p Parser
	p.s = make([]string, 0)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command: ")
		input, _ := reader.ReadString('\n')
		if input[:len(input)-1] == "exit" {
			break
		}
		if input[len(input)-2] == ';' {
			p.s = append(p.s, input[:len(input)-2])
			fmt.Print("	>")
		} else {
			p.s = append(p.s, input[:len(input)-1])
			executeCommands(p.s)
			p.s = nil
		}
	}
}
