package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	ans := 0
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for _, v := range g {
		pos := sort.Search(len(s), func(i int) bool {
			return s[i] >= v
		})
		if pos != len(s) {
			ans++
			s = append(s[:pos], s[pos+1:]...)
		}
		if len(s) == 0 {
			return ans
		}
	}
	return ans
}
func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
