package main

import "fmt"

func subsets(nums []int) (res [][]int) {
	res = append(res, []int{})
	for _, val := range nums {
		for _, v := range res {
			tmp := make([]int, len(v))
			copy(tmp, v)
			tmp = append(tmp, val)
			res = append(res, tmp)
		}
	}
	return
}
func main() {
	//res := make([][]int, 0)
	//tmp := res
	//tmp = append(tmp, []int{1})
	//fmt.Println(res)
	//fmt.Println(tmp)
	//fmt.Println(subsets([]int{1, 2, 3}))
	//fmt.Println(subsets([]int{0}))
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}
