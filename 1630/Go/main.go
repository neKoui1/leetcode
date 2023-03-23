package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) (res []bool) {
	m := len(l)
	n := len(r)
	if n != m {
		return nil
	}
	for i := 0; i < m; i++ {
		temp := make([]int, r[i]-l[i])
		copy(temp, nums[l[i]:r[i]])
		temp = append(temp, nums[r[i]])
		sort.SliceStable(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		dif := temp[1] - temp[0]
		flag := true
		for i := 0; i < len(temp)-1 && flag; i++ {
			if temp[i+1]-temp[i] != dif {
				flag = false
			}
		}
		res = append(res, flag)

	}
	return
}
func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	fmt.Println(checkArithmeticSubarrays([]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10},
		[]int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}))
}
