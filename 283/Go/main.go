package main

import "fmt"

func moveZeroes(nums []int) {
	cnt := 0
	for _, v := range nums {
		if v != 0 {
			nums[cnt] = v
			cnt++
		}
	}
	for i := cnt; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
	nums2 := []int{0}
	moveZeroes(nums2)
	fmt.Println(nums2)
}
