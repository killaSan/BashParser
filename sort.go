package main

import ()

// Sort command: sort [-r] file [...]

type Sort struct {
	flag  string
	files []string
}

func merge(l, r []string, comp func(a, b string) bool) []string {
	var i, j, k int
	res := make([]string, len(l)+len(r))
	for i < len(l) && j < len(r) {
		if comp(l[i], r[j]) {
			res[k] = l[i]
			i += 1
			k += 1
		} else {
			res[k] = r[j]
			k += 1
			j += 1
		}
	}
	for i < len(l) {
		res[k] = l[i]
		k += 1
		i += 1
	}
	for j < len(r) {
		res[k] = r[j]
		k += 1
		j += 1
	}
	return res
}

func mergeSort(arr []string, comp func(a, b string) bool) []string {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(mergeSort(left, comp), mergeSort(right, comp), comp)
}

func (s Sort) executeCommand() []string {
	var res []string
	var resFromMergeSort []string
	for i, _ := range s.files {
		if exists(s.files[i]) {
			lines, _, _, _ := wc(s.files[i])
			tempArr := removeAllTabs(head(lines, s.files[i]))
			if s.flag == "r" {
				resFromMergeSort = mergeSort(tempArr, func(a, b string) bool { return a > b }) // descending
			} else {
				resFromMergeSort = mergeSort(tempArr, func(a, b string) bool { return a < b }) // ascending
			}
			res = append(res, resFromMergeSort...)
		}
	}
	return res
}

func (s Sort) readCommand(input string) (string, []string) {
	remain := input[5:]
	if remain[0:2] == "-r" {
		s.flag = "r"
		remain = remain[3:]
	}
	s.files = divideFiles(remain)
	return s.flag, s.files
}

func (s Sort) startCommand(str string) {
	s.flag, s.files = s.readCommand(str)
	res := s.executeCommand()
	printSlice(res)
}

func removeAllTabs(s []string) []string {
	for i, _ := range s {
		s[i] = removeTabs(s[i])
	}
	return s
}
