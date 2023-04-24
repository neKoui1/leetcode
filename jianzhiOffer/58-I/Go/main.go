package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var res string
	var i = len(s) - 1
	var j = i
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}
		res = res + s[i+1:j+1] + " "
	}
	return strings.TrimRight(res, " ")
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords(" hello world!   "))
	fmt.Println(reverseWords("a good     example"))
}
