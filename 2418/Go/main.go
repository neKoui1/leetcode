package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	ans := make([]string, 0)
	h2name := make(map[int]string)
	n := len(names)
	for i := 0; i < n; i++ {
		h2name[heights[i]] = names[i]
	}
	sort.SliceStable(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	for i := 0; i < n; i++ {
		ans = append(ans, h2name[heights[i]])
	}
	return ans
}
func main() {

}
