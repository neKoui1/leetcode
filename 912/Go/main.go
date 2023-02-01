package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
