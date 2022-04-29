package main

import "fmt"

func findTheDifference(s string, t string) byte {
	hashs := make(map[rune]int)
	hasht := make(map[rune]int)
	for _, v := range s {
		hashs[v]++
	}
	for _, v := range t {
		hasht[v]++
	}
	for _, v := range t {
		val, flag := hashs[v]
		if !flag {
			return byte(v)
		} else {
			if val != hasht[v] {
				return byte(v)
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
	fmt.Println(findTheDifference("", "y"))
	fmt.Println(findTheDifference("a", "aa"))
}
