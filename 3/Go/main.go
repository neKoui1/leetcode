package main

import "fmt"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func lengthOfLongestSubstring(s string) int {
	start, end, length, maxLen := 0, 0, len(s), 1
	if length == 0 {
		return 0
	}
	for end < length {
		temp := start
		for i := temp; i < end; i++ {
			if s[i] == s[end] {
				start = i + 1
			}
		}
		maxLen = Max(end-start+1, maxLen)
		end++
	}
	return maxLen
}
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
