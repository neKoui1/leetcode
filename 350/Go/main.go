package main

import "fmt"

func intersect(nums1 []int, nums2 []int) (res []int) {
	hash := make(map[int]int)
	for _, v := range nums1 {
		hash[v]++
	}

	for _, v := range nums2 {
		if _, ok := hash[v]; ok && hash[v] > 0 {
			res = append(res, v)
			hash[v]--
		}
	}
	return
}
func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
