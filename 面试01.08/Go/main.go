package main

func setZeroes(matrix [][]int) {
	rowIndex := make([]int, 0)
	colIndex := make([]int, 0)
	for i, val := range matrix {
		for j, v := range val {
			if v == 0 {
				rowIndex = append(rowIndex, i)
				colIndex = append(colIndex, j)
			}
		}
	}
	for _, v := range rowIndex {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v][i] = 0
		}
	}
	for _, v := range colIndex {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}
}
func main() {
	setZeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	})
}
