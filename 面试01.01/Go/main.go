package main

import (
	"fmt"
	"strings"
)

func isUnique(astr string) bool {
	for _, v := range astr {
		if strings.Count(astr, string(v)) != 1 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isUnique("leetcode"))
	fmt.Println(isUnique("abc"))
}
