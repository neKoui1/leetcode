package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) (ans [][]string) {
	x := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.SliceStable(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		temp := string(s)
		x[temp] = append(x[temp], str)
	}
	for _, v := range x {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
}
