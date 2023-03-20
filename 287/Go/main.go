package main

import "fmt"

/**
带环链表
快慢指针找出带环链表的入口
*/
func findDuplicate(nums []int) int {
	slow := nums[nums[0]]
	fast := nums[nums[nums[0]]]
	for slow != fast {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
