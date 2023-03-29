package main

import "fmt"

func reverse(s []byte) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

func reverseStr(s string, k int) string {
	res := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		if i+k < len(s) {
			reverse(res[i : i+k])
		} else {
			reverse(res[i:len(s)])
		}
	}
	return string(res)
}
func main() {
	fmt.Println(reverseStr("abcdefg", 2))
	fmt.Println(reverseStr("abcd", 2))
}
