package main

import (
	"fmt"
	"reflect"
)

func CheckPermutation(s1 string, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		return false
	}
	temps1 := make(map[byte]int)
	temps2 := make(map[byte]int)
	for i := 0; i < l1; i++ {
		temps1[s1[i]]++
		temps2[s2[i]]++
	}
	return reflect.DeepEqual(temps1, temps2)
}
func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
	fmt.Println(CheckPermutation("abc", "bad"))
}
