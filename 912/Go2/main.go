package main

import "fmt"

func QuickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	x, i, j := nums[(l+r)>>1], l-1, r+1
	for i < j {
		i++
		for nums[i] < x {
			i++
		}
		j--
		for nums[j] > x {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	QuickSort(nums, l, j)
	QuickSort(nums, j+1, r)
}

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}
func main() {
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}
