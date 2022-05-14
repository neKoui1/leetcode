package main

import (
	"fmt"
	"strings"
)

type Set map[rune]int

func findWords(words []string) (res []string) {
	set := make(Set)
	for _, v := range "qwertyuiop" {
		set[v] = 1
	}
	for _, v := range "asdfghjkl" {
		set[v] = 2
	}
	for _, v := range "zxcvbnm" {
		set[v] = 3
	}
	temp := make([]string, 0)
	for _, v := range words {
		temp = append(temp, strings.ToLower(v))
	}
	for i, v := range temp {
		flag := true
		for i := 1; i < len(v) && flag; i++ {
			if set[rune(v[i])] != set[rune(v[i-1])] {
				flag = false
			}
		}
		if flag {
			res = append(res, words[i])
		}
	}
	return
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}
