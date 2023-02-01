package main

import "fmt"

/*
	暴力解法
	会超时
*/

//func minSubArrayLen(target int, nums []int) int {
//	sum := make([]int, 0)
//	sum = append(sum, nums[0])
//	min := 0x7fffffff
//	flag := false
//	for i := 1; i < len(nums); i++ {
//		sum = append(sum, sum[i-1]+nums[i])
//	}
//	if sum[len(sum)-1] >= target && len(sum) < min {
//		flag = true
//		min = len(sum)
//	}
//	for i := 0; i < len(sum)-1; i++ {
//		if sum[i] >= target && i+1 < min {
//			min = i + 1
//		}
//		for j := i + 1; j < len(sum); j++ {
//			if sum[j]-sum[i] >= target && j-i < min {
//				flag = true
//				min = j - i
//			}
//		}
//	}
//	if flag {
//		return min
//	}
//	return 0
//}

/*
	滑动窗口解法
*/
func minSubArrayLen(target int, nums []int) int {
	start, end, sum, min := 0, 0, 0, 0x7fffffff
	length := len(nums)
	for end < length {
		sum += nums[end]
		for sum >= target {
			if end-start+1 < min {
				min = end - start + 1
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if min == 0x7fffffff {
		return 0
	}
	return min
}
func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(5, []int{2, 3, 1, 1, 1, 1, 1}))
}
