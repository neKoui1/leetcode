package main

import "fmt"

func equalFrequency(word string) bool {
	charCount := [26]int{}
	for _, c := range word {
		charCount[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if charCount[i] == 0 {
			continue
		}
		charCount[i]--
		frequency := make(map[int]bool)
		for _, f := range charCount {
			if f > 0 {
				frequency[f] = true
			}
		}
		if len(frequency) == 1 {
			return true
		}
		charCount[i]++
	}
	return false
}
func main() {
	fmt.Println(equalFrequency("abcc"))
	fmt.Println(equalFrequency("aazz"))
}
