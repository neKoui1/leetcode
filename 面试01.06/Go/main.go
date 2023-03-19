package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	cnt := 1
	res := ""
	ch := S[0]
	for i := 1; i < len(S); i++ {
		if S[i] == ch {
			cnt++
		} else {
			res += string(ch) + strconv.Itoa(cnt)
			cnt = 1
			ch = S[i]
		}
	}
	res += string(ch) + strconv.Itoa(cnt)
	if len(res) >= len(S) {
		return S
	}
	return res
}
func main() {
	fmt.Println(compressString("aabccccaaa"))
	fmt.Println(compressString("abbccd"))
}
