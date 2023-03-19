package main

import (
	"fmt"
	"strings"
)

/**
Obviously, KMP is more suitable here than violence
*/
func isFlipedString(s1 string, s2 string) bool {
	//if len(s1) != len(s2) {
	//	return false
	//}
	//if s1 == "" && s2 == "" {
	//	return true
	//}
	//temp := s2 + s2
	//for i := 0; i < len(temp)-len(s1); i++ {
	//	if s1 == temp[i:i+len(s1)] {
	//		return true
	//	}
	//}
	//return false

	/**
	But go provides strings pkg..
	So I can write like this
	STRANGE KNOWLEDGE ADDED
	*/
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}
func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
	fmt.Println(isFlipedString("aa", "aba"))
}
