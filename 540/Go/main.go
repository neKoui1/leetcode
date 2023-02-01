package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}

	return res
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}
