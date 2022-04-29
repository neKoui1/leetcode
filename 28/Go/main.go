package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	l := len(needle)
	for i := range haystack {
		if len(haystack)-i < l {
			break
		}
		if haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("aaa", "aaaa"))
}
