package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	cal := func(a, b, c float64) float64 {
		p := (a + b + c) / 2
		return math.Sqrt(p * (p - a) * (p - b) * (p - c))
	}
	dis := func(x1, y1, x2, y2 int) float64 {
		return math.Sqrt(float64(x1-x2)*float64(x1-x2) + float64(y1-y2)*float64(y1-y2))
	}
	var max float64 = 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				cur := cal(dis(points[i][0], points[i][1], points[j][0], points[j][1]), dis(points[i][0], points[i][1], points[k][0], points[k][1]), dis(points[k][0], points[k][1], points[j][0], points[j][1]))
				if cur > max {
					max = cur
				}
			}
		}
	}
	return max
}
func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}
