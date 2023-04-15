package main

import "fmt"

func camelMatch(queries []string, pattern string) (res []bool) {
	for _, v := range queries {
		res = append(res, isMatch(v, pattern))
	}
	return
}

func isMatch(query string, pattern string) bool {
	index := 0
	for _, v := range query {
		if index == len(pattern) {
			if v > 'z' || v < 'a' {
				return false
			}
		} else {
			if v == int32(pattern[index]) {
				index++
			} else if v > 'z' || v < 'a' {
				return false
			}
		}
	}
	return index == len(pattern)
}
func main() {
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))
}
