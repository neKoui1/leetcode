package main

import (
	"fmt"
	"reflect"
)

func firstUniqChar(s string) byte {
	if reflect.DeepEqual(s, "") {
		return ' '
	}
	kv := make(map[int32]int)
	for _, v := range s {
		kv[v-'a']++
	}
	for _, v := range s {
		if kv[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
	fmt.Println(firstUniqChar(""))
}
