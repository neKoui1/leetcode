package main

import "fmt"

func oneEditAway(first string, second string) bool {
	m := len(first)
	n := len(second)

	if n > m {
		return oneEditAway(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, v := range second {
		if first[i] != byte(v) {
			if m == n {
				return first[i+1:] == second[i+1:]
			} else {
				return first[i+1:] == second[i:]
			}
		}
	}
	/*
		遍历结束，两个字符串相等故返回true
	*/
	return true
}
func main() {
	fmt.Println(oneEditAway("pale", "ple"))
	fmt.Println(oneEditAway("pales", "pal"))
}
