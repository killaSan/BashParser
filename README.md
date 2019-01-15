# BashParser

INTRODUCTION

The program is a Bash parser which will perform the following commands
	- head [with flags -n - number of lines to be outputted][file - passed to head]...
	- cat
	- sort
	- tail
	- wc [-c (counts characters)/ -w(counts words)/ -l(counts lines)/ -L(length of the longest line per file) ] [file - passed to head ]...
	- cut
	- grep
	- cmp
	- cp
	- mv
	- rm

INSTALLING
	1) go get github.com/killaSan/helperFunctions
	2) github.com/killaSan/BashParser
	3) go run interface.go 

RUNNING
	The input will be a Bash command. For instance - head file.txt/ head -100 book.txt example.go etc. If a wrong input is passed, the message "Wrong command!" will be printed, followed by another input by the user. If exit is called the program will stop its work.

TESTS 
	// TODO
