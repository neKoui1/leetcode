package main

import "fmt"

// 不单调
func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + ((right - left) >> 1)
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return numbers[left]
}
func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
