package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) (cnt int) {
	if k == 0 {
		return
	}
	N := len(nums)
	if N == 1 {
		if nums[0] < k {
			cnt++
		}
		return
	}
	mult, p, q := 1, 0, 1
	for p != N {
		mult *= nums[p]
		if mult < k {
			cnt++
		}
		for q != N {
			mult *= nums[q]
			q++
			if mult < k {
				cnt++
			} else if mult >= k {
				break
			}
		}
		mult = 1
		p++
		q = p + 1
	}
	return
}
func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}
