package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Ints(arr)
	sort.SliceStable(arr, func(i, j int) bool {
		si := fmt.Sprintf("%b", arr[i])
		sj := fmt.Sprintf("%b", arr[j])
		cnti, cntj := 0, 0
		for _, v := range si {
			if v-'0' == 1 {
				cnti++
			}
		}
		for _, v := range sj {
			if v == '1' {
				cntj++
			}
		}
		return cnti < cntj
	})
	return arr
}

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}
