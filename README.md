# BashParser

INTRODUCTION

The program is a Bash parser which will perform the following commands
	- head [-n] files[...] ;
	- cat file [...];
	-  ort [-r] file [...];
	- tail [-n] files [...];
	- wc [-c (counts characters)/ -w(counts words)/ -l(counts lines)/ ] [file - passed to head ]... ;
	- cut -c beg-end files[...] ;
	- grep PATTERN files [...] ;
	- cp src dst;
	- mv src dst;
	- rm files [...];

INSTALLING
	1) go get github.com/killaSan/BashParser
	2) go run interface.go 

RUNNING
	The input will be a Bash command. For instance - head file.txt/ head -100 book.txt example.go etc. If a wrong input is passed, the message "Wrong command!" will be printed, followed by another input by the user. If exit is called the program will stop its work.
