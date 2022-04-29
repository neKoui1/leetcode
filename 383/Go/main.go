package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	hashm := make(map[rune]int)
	hashr := make(map[rune]int)
	for _, v := range magazine {
		hashm[v]++
	}
	for _, v := range ransomNote {
		hashr[v]++
	}

	flag := true
	for _, v := range ransomNote {
		if _, ok := hashm[v]; !ok {
			flag = false
			break
		}
		if hashr[v] > hashm[v] {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
