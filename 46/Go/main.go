package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	} else {
		res := make([][]int, 0)
		for frequency := len(nums); frequency > 0; frequency-- {
			tmp := permute(nums[1:])
			for _, v := range tmp {
				r := append([]int{nums[0]}, v...)
				res = append(res, r)
			}
			tmp = append(tmp[1:], tmp[0])
			nums = append(nums[1:], nums[0])
		}
		return res
	}
}
func main() {
	fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{1, 0}))
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(len(permute([]int{1, 2, 3, 4, 5})))
}
