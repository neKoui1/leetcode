package main

import "fmt"

func maxPower(s string) int {
	max, cnt := 1, 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		}
		if cnt > max {
			max = cnt
		}
		if s[i] != s[i+1] {
			cnt = 1
		}
	}
	return max
}
func main() {
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
}
