package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func lengthOfLongestSubstring(s string) (ans int) {
	left := 0
	set := make([]int, 128)
	for right, c := range s {
		set[c]++
		for set[c] > 1 {
			set[s[left]]--
			left++
		}
		ans = Max(ans, right-left+1)
	}
	return
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
