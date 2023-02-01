package main

import "fmt"

/*
	超时
*/
//func numSubarraysWithSum(nums []int, goal int) int {
//	sum := make([]int, len(nums))
//	cnt := 0
//	sum[0] = nums[0]
//	for i := 1; i < len(nums); i++ {
//		sum[i] = nums[i] + sum[i-1]
//	}
//	fmt.Println(sum)
//	for i := 0; i < len(sum); i++ {
//		if sum[i] == goal {
//			cnt++
//		}
//	}
//	for i := 0; i < len(sum)-1; i++ {
//		for j := i + 1; j < len(sum); j++ {
//			if sum[j]-sum[i] == goal {
//				cnt++
//			}
//		}
//	}
//	return cnt
//}

/*
	所以试试哈希表
*/
func numSubarraysWithSum(nums []int, goal int) int {
	sum := make([]int, len(nums))
	cnt := 0
	pair := make(map[int]int)
	sum[0] = nums[0]
	pair[sum[0]]++
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		pair[sum[i]]++
	}
	if _, ok := pair[0]; ok && goal == 0 {
		for i := 1; i <= pair[0]; i++ {
			cnt += i
		}
	}
	for i, v := range pair {
		if i < goal {
			continue
		} else if i == goal && i != 0 && goal != 0 {
			cnt += v
		}
		if value, ok := pair[i-goal]; ok && i != 0 && goal != 0 {
			cnt += value * v
			continue
		}
		if value, ok := pair[i-goal]; ok && i != 0 && goal == 0 {
			for j := 1; j < value; j++ {
				cnt += j
			}
		}
	}
	return cnt
}
func main() {
	fmt.Println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 1}, 2))
	fmt.Println(numSubarraysWithSum([]int{0, 1, 1, 1, 1}, 3))
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}, 0))
}
