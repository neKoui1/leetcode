package main

import (
	"fmt"
)

func canPermutePalindrome(s string) bool {
	hash := make(map[int32]int)
	for _, v := range s {
		hash[v]++
	}
	cnt := 0
	if len(hash) == 1 {
		return true
	}
	for _, v := range hash {
		if v%2 != 0 {
			cnt++
			if cnt == 2 {
				return false
			}
		} else {
			continue
		}
	}
	return true
}
func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
	fmt.Println(canPermutePalindrome("aaa"))
}
