package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	copy := strings.TrimSpace(s)
	count := 1
	if s == " " {
		return 0
	}
	for i, v := range copy {
		if v == ' ' {
			count = len(copy[i+1:])
		}
	}
	cnt := 0
	for _, v := range copy {
		if v != ' ' {
			cnt++
		}
	}
	if cnt == len(copy) {
		return len(copy)
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord(""))
}
