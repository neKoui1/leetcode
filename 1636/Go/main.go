package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key int
	Cnt int
}

func frequencySort(nums []int) []int {
	temp := make(map[int]int, 0)
	for _, v := range nums {
		temp[v]++
	}
	tempArr := make([]Pair, 0)
	res := make([]int, 0)
	for i, v := range temp {
		tempArr = append(tempArr, Pair{i, v})
	}
	sort.SliceStable(tempArr, func(i, j int) bool {
		if tempArr[i].Cnt == tempArr[j].Cnt {
			return tempArr[i].Key > tempArr[j].Key
		}
		return tempArr[i].Cnt < tempArr[j].Cnt
	})
	for _, v := range tempArr {
		for k := 0; k < v.Cnt; k++ {
			res = append(res, v.Key)
		}
	}
	return res
}

func main() {
	fmt.Println(frequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(frequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
