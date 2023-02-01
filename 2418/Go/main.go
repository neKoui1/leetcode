package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	nh := make(map[int]string, 0)
	for i := 0; i < len(names); i++ {
		nh[heights[i]] = names[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	res := make([]string, 0)
	for i := 0; i < len(heights); i++ {
		res = append(res, nh[heights[i]])
	}
	return res
}

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}
