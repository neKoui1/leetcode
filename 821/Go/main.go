package main

import "fmt"

func Abs(i, j int) int {
	if i > j {
		return i - j
	} else {
		return j - i
	}
}

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	for i, v := range s {
		if v == rune(c) {
			ans[i] = 0
		}
	}
	for i := range s {
		for j, v := range s {
			if v == rune(c) {
				if Abs(i, j) < ans[i] {
					ans[i] = Abs(i, j)
				}
			}
		}
	}

	return ans
}
func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
