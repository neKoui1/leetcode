package main

import "fmt"

/**
hash
空间复杂度O(m+n)
时间复杂度O(n)
*/

//func intersection(nums1 []int, nums2 []int) (res []int) {
//	hash1 := make(map[int]int)
//	hash2 := make(map[int]int)
//	for _, v := range nums1 {
//		hash1[v]++
//	}
//	for _, v := range nums2 {
//		hash2[v]++
//	}
//	for i := range hash1 {
//		if _, ok := hash2[i]; ok {
//			res = append(res, i)
//		}
//	}
//	return
//}

/**
two pointers
*/
func intersection(nums1 []int, nums2 []int) (res []int) {

	return
}
func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
