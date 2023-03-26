package main

import (
	"fmt"
	"math"
)

func judge(a int) bool {
	if a == 0 || a == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func primeSubOperation(nums []int) bool {
	n := len(nums)

LOOP:
	for i := 0; i < n-1; i++ {

		if nums[i] >= nums[i+1] {
			sub := nums[i] - nums[i+1]

			for j := nums[i] - 1; j >= sub; j-- {
				if judge(j) {
					nums[i] -= j
					goto LOOP
				}
			}
		}
	}

	for i := 0; i < n-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(primeSubOperation([]int{4, 9, 6, 10}))
	fmt.Println(primeSubOperation([]int{6, 8, 11, 12}))
	fmt.Println(primeSubOperation([]int{5, 8, 3}))
}
