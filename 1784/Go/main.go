package main

import "fmt"

func checkOnesSegment(s string) bool {
	find1, find2 := 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '1' {
			find2 = i
			break
		}
	}
	for i := find1; i <= find2; i++ {
		if s[i] == '0' {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(checkOnesSegment("1001"))
	fmt.Println(checkOnesSegment("110"))
}
