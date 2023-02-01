package main

import "fmt"

func judge(hash map[int32]int, word string) bool {
	for _, v := range word {
		if _, ok := hash[v]; !ok {
			return false
		}
	}
	return true
}
func countConsistentStrings(allowed string, words []string) (res int) {
	hash := make(map[int32]int)
	for _, v := range allowed {
		hash[v]++
	}
	for _, v := range words {
		if judge(hash, v) {
			res++
		} else {
			continue
		}
	}
	return res
}
func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
