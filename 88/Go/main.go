package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Println(nums1)
}
