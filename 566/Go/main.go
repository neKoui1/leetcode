package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	temp := make([]int, 0)
	for _, v := range mat {
		temp = append(temp, v...)
	}
	res := make([][]int, 0)
	line := make([]int, 0)
	cnt := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			line = append(line, temp[cnt])
			cnt++
		}
		res = append(res, line)
		line = nil
	}
	return res
}

func main() {
	mat := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrixReshape(mat, 1, 4))
	fmt.Println(matrixReshape(mat, 2, 4))
	fmt.Println(matrixReshape(mat, 4, 1))
	mat = [][]int{
		{1, 2, 3, 4},
	}
	fmt.Println(matrixReshape(mat, 2, 2))
}
