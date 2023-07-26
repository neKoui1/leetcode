package main

import (
	"fmt"
	"reflect"
)

func getArr(s string) (ans []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		ans = append(ans, int(s[i+1]-s[i]))
	}
	return ans
}

func oddString(words []string) string {
	diff1 := getArr(words[0])
	diff2 := getArr(words[1])
	n := len(words)
	flag := reflect.DeepEqual(diff1, diff2)
	for i := 2; i < n; i++ {
		diff := getArr(words[i])
		if !reflect.DeepEqual(diff, diff1) && flag {
			return words[i]
		} else if !flag && reflect.DeepEqual(diff, diff1) {
			return words[1]
		} else if !flag && reflect.DeepEqual(diff, diff2) {
			return words[0]
		}
	}
	return ""
}
func main() {
	fmt.Println(oddString([]string{"adc", "wzy", "abc"}))
	fmt.Println(oddString([]string{"aaa", "bob", "ccc", "ddd"}))
}
