package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	flag := true
	if coordinates[0][1] == coordinates[1][1] {
		for i := 0; i < len(coordinates)-1; i++ {
			if coordinates[i][1] != coordinates[i+1][1] {
				flag = false
			}
		}
	} else {
		k := float64(coordinates[0][0]-coordinates[1][0]) / float64(coordinates[0][1]-coordinates[1][1])
		for i := 0; i < len(coordinates)-1; i++ {
			if k != float64(coordinates[i][0]-coordinates[i+1][0])/float64(coordinates[i][1]-coordinates[i+1][1]) {
				flag = false
			}
		}
	}

	return flag
}

func main() {
	coordinates := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	fmt.Println(checkStraightLine(coordinates))
}
