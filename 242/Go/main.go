package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash1 := make(map[byte]int)
	hash2 := make(map[byte]int)
	for _, v := range s {
		hash1[byte(v)]++
	}
	for _, v := range t {
		hash2[byte(v)]++
	}
	for _, v := range s {
		if hash1[byte(v)] != hash2[byte(v)] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "cat"))
	fmt.Println(isAnagram("a", "ab"))
}
