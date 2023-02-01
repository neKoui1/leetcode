package main

import "fmt"

func minOrMax(i, a, b int) int {
	if i%2 == 0 {
		if a > b {
			return b
		}
		return a
	} else {
		if a > b {
			return a
		}
		return b
	}
}
func minMaxGame(nums []int) (res int) {
	n := len(nums)

	for n != 1 {
		newNums := make([]int, n/2)
		n = len(newNums)
		for i := 0; i < n; i++ {
			newNums[i] = minOrMax(i, nums[2*i], nums[2*i+1])
		}
		nums = newNums
	}
	return nums[0]
}

func main() {
	fmt.Println(minMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2}))
	fmt.Println(minMaxGame([]int{3}))
}
