package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if j == len(nums2)-1 {
					ans = append(ans, -1)
				} else {
					for k := j + 1; k < len(nums2); k++ {
						if nums2[k] > nums2[j] {
							ans = append(ans, nums2[k])
							break
						}
						if k == len(nums2)-1 {
							ans = append(ans, -1)
						}
					}
				}
			}
		}
	}
	return ans
}

func main() {
	nums1 := []int{1, 3, 5, 2, 4}
	nums2 := []int{6, 5, 4, 3, 2, 1, 7}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
