package main

import "fmt"

func checkIfPangram(sentence string) bool {
	has := make(map[int32]int)
	for _, v := range sentence {
		has[v]++
	}
	if len(has) != 26 {
		return false
	}
	return true
}
func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}
