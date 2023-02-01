package main

import (
	"fmt"
)

func decodeMessage(key string, message string) string {
	dict := make(map[string]byte)
	cnt := 0
	for _, v := range key {
		if v == ' ' {
			continue
		}
		if _, ok := dict[string(v)]; ok {
			continue
		}
		dict[string(v)] = byte(cnt + 'a')
		cnt++
	}
	res := make([]byte, 0)
	for _, v := range message {
		if v == ' ' {
			res = append(res, ' ')
			continue
		}
		res = append(res, dict[string(v)])
	}
	return string(res)
}
func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}
