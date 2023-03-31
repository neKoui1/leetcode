package main

import (
	"fmt"
	"sort"
)

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func numberOfWeakCharacters(properties [][]int) (res int) {
	sort.SliceStable(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	n := len(properties)
	maxDef := properties[0][1]
	for i := 1; i < n; i++ {
		maxDef = Max(maxDef, properties[i][1])
		if properties[i][1] < maxDef {
			res++
		}
	}

	return
}
func main() {
	fmt.Println(numberOfWeakCharacters([][]int{
		{5, 5},
		{6, 3},
		{3, 6},
	}))
	fmt.Println(numberOfWeakCharacters([][]int{
		{2, 2},
		{3, 3},
	}))
	fmt.Println(numberOfWeakCharacters([][]int{
		{1, 5},
		{10, 4},
		{4, 3},
	}))
	fmt.Println(numberOfWeakCharacters([][]int{
		{2, 3},
		{2, 2},
		{3, 2},
	}))
}
