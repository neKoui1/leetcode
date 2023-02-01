package main

import "fmt"

/*
	比较暴力的滑动窗口
	会超时
*/
//func containsNearbyDuplicate(nums []int, k int) bool {
//	flag, start, end, length := false, 0, 0, len(nums)
//	for end < length {
//		for i := start; i < end; i++ {
//			if nums[i] == nums[end] {
//				flag = true
//				if end-i <= k {
//					return true
//				}
//			}
//			if flag && i == end-1 {
//				start++
//			}
//		}
//		end++
//	}
//	return false
//}
/*
	用集合来提前处理掉不符合abs比较的解法
	依然是滑动窗口 不过这种写法很强
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{}, 0)
	for i, v := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
