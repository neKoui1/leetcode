package main

import "fmt"

func diagonalSum(mat [][]int) int {
	sum := 0
	for i, val := range mat {
		for j, v := range val {
			if i == j || i+j == len(mat)-1 {
				sum += v
			}
		}
	}

	return sum
}

func main() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(diagonalSum(mat))
	mat = [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}
	fmt.Println(diagonalSum(mat))
	mat = [][]int{
		{5},
	}
	fmt.Println(diagonalSum(mat))
}
