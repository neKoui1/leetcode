package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key rune
	Cnt int
}

func frequencySort(s string) string {
	temp := make(map[rune]int, 0)
	for _, v := range s {
		temp[v]++
	}
	arr := make([]Pair, 0)
	for i, v := range temp {
		arr = append(arr, Pair{
			i,
			v,
		})
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].Cnt > arr[j].Cnt
	})
	res := ""
	for _, v := range arr {
		for j := 0; j < v.Cnt; j++ {
			res += string(v.Key)
		}
	}
	return res
}
func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}
