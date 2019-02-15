package main

import (
	"fmt"
)

/*
	Type Parser are the inputs given by the user, which are later turned into commands
*/
type Parser struct {
	s []string
}


/*
	Command interface: Types Head, Mv, Cat, Copy, Cut, Grep, Rm, Sort, Tail and Wc implement it
readCommand - transfers input from user into Bash Command
executeCommand - executes the command for each given file to the command
startCommand - executes both readCommand and executeCommand
*/

type Command interface {
	startCommand(string)
	executeCommand() []string
	readCommand(string) (string, []string)
}

func executeCommands(p []string) {
	for i, _ := range p {
		if len(p[i]) == 4 && p[i] == "exit" {
			break
		} else if len(p[i]) >= 5 && p[i][:5] == "head " && countWords(p[i]) >= 3 {
			var h Head
			h.startCommand(p[i])
		} else if len(p[i]) >= 3 && p[i][:3] == "mv " && countWords(p[i]) >= 2 {
			var m Mv
			m.startCommand(p[i])
		} else if len(p[i]) >= 4 && p[i][:4] == "cat " && countWords(p[i]) >= 2 {
			var c Cat
			c.startCommand(p[i])
		} else if len(p[i]) >= 3 && p[i][:3] == "cp " && countWords(p[i]) >= 2 {
			var c Copy
			c.startCommand(p[i])
		} else if len(p[i]) >= 4 && p[i][:4] == "cut " && countWords(p[i]) >= 4 {
			var c Cut
			c.startCommand(p[i])
		} else if len(p[i]) >= 5 && p[i][:5] == "grep " && countWords(p[i]) >= 3 {
			var g Grep
			g.startCommand(p[i])
		} else if len(p[i]) >= 3 && p[i][:3] == "rm " && countWords(p[i]) >= 2 {
			var r Rm
			r.startCommand(p[i])
		} else if len(p[i]) >= 5 && p[i][:5] == "sort " && countWords(p[i]) >= 2 {
			var s Sort
			s.startCommand(p[i])
		} else if len(p[i]) >= 5 && p[i][:5] == "tail " && countWords(p[i]) >= 2 {
			var t Tail
			t.startCommand(p[i])
		} else if len(p[i]) >= 3 && p[i][:3] == "wc " && countWords(p[i]) >= 2 {
			var w Wc
			w.startCommand(p[i])
		} else {
			fmt.Println("Wrong command!")
		}
		fmt.Println()
		fmt.Println("-----------")
	}
}
