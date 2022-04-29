package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	s := fmt.Sprintf("%b", n)
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(699050))
}
