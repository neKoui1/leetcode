package main

import "sort"

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	res := -1
	for i := 1; i < len(points); i++ {
		res = Max(res, points[i][0]-points[i-1][0])
	}
	return res
}

func main() {

}
