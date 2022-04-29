package main

import "fmt"

func Abs(x1, x2 int) int {
	if x1 > x2 {
		return x1 - x2
	} else {
		return x2 - x1
	}
}

func ManhattanDis(x1, y1 int, x2, y2 int) int {
	return Abs(x1, x2) + Abs(y1, y2)
}

func nearestValidPoint(x int, y int, points [][]int) int {
	distance := 0xffffffff
	index := -1
	for i := 0; i < len(points); i++ {
		if points[i][0] == x {
			if ManhattanDis(x, y, points[i][0], points[i][1]) < distance {
				distance = ManhattanDis(x, y, points[i][0], points[i][1])
				index = i
			}
		}
		if points[i][1] == y {
			if ManhattanDis(x, y, points[i][0], points[i][1]) < distance {
				distance = ManhattanDis(x, y, points[i][0], points[i][1])
				index = i
			}
		}
	}

	return index
}

func main() {
	points := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	}
	fmt.Println(nearestValidPoint(3, 4, points))
}
