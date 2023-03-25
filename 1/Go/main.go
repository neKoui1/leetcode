package main

import (
	"fmt"
)

func twoSum(nums []int, target int) (res []int) {
	hash := make(map[int]int)
	for i, v := range nums {
		hash[v] = i
	}
	for i, v := range nums {
		if val, ok := hash[target-v]; ok && val != i {
			return []int{i, val}
		}
	}
	return
}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
